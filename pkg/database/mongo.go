package database

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/gertd/go-pluralize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 注册表和索引
// 标签 mongo
// 索引 index
// 复合索引 compound
// 示例
// type Order struct {
// 	Id          primitive.ObjectID `bson:"_id" mongo:"compound:'created_time_id_appid'"`
// 	Price       int                `bson:"price"`
// 	Appid       string             `bson:"appid" mongo:"compound:'created_time_id_appid, appid_refund_date, appid_pay_date'"` // application appid
// 	Title       string             `bson:"title"`
// 	TheaterId   int                `bson:"theater_id"`
// 	OrderId     string             `bson:"order_id" mongo:"index"`
// 	Channel     string             `bson:"channel"` // douyin/wechat/kuaishou
// 	Status      int                `bson:"status"`
// 	CreatedTime string             `bson:"created_time" mongo:"index; compound:'created_time_id_appid'"`
// 	PayDate     string             `bson:"pay_date" mongo:"index; compound:'appid_pay_date'"`
// 	RefundDate  string             `bson:"refund_date" mongo:"compound:'appid_refund_date'"`
// 	SourceId    string             `bson:"source_id"`
// 	SourceType  string             `bson:"source_type"`
// 	MemberId    string             `bson:"member_id"`
// 	CreatedAt   string             `bson:"created_at" mongo:"index"` // 添加时间 YYYY-MM-dd hh:mm:ss
// 	UpdatedAt   string             `bson:"updated_at"`               // 更新时间 YYYY-MM-dd hh:mm:ss
// }

func AutoRegisterMongo(db *mongo.Database, v ...interface{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, val := range v {
		t := reflect.TypeOf(val)

		p := pluralize.NewClient()
		plural := p.Plural(t.Name())
		colName := strings.ToLower(plural[:1] + plural[1:])
		indexList := []string{}
		compoundMap := make(map[string]map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			f := field.Tag.Get("bson")
			if f == "" {
				LogPanic("Mongo", "字段标签bson不能为空")
			}
			mctx := field.Tag.Get("mongo")
			if mctx != "" {
				// 处理mongo定义索引
				// 切分
				slice := []string{}
				for _, v := range strings.Split(mctx, ";") {
					slice = append(slice, strings.TrimSpace(v))
				}

				for _, v := range slice {
					// 独立索引
					if v == "index" {
						indexList = append(indexList, f)
					}

					// 联合索引
					if strings.Contains(v, "compound") {

						comp := strings.ReplaceAll(v, "compound:", "")

						var compList []string

						for _, val := range strings.Split(comp, ",") {
							compList = append(compList, strings.ReplaceAll(strings.TrimSpace(val), "'", ""))
						}

						for _, compName := range compList {
							if compoundMap[compName] == nil {
								compoundMap[compName] = make(map[string]interface{})
							}
							compoundMap[compName][f] = nil
						}
					}
				}
			}
		}

		// 校验联合索引长度
		// 小于2则panic
		var indexModelList []mongo.IndexModel
		for _, compItem := range compoundMap {
			if len(compItem) < 2 {
				LogPanic("Mongo", "联合索引字段不足")
			}
			keys := bson.D{}
			for indexItem := range compItem {
				keys = append(keys, bson.E{Key: indexItem, Value: 1})
			}
			indexModelList = append(indexModelList, mongo.IndexModel{Keys: keys, Options: nil})
		}

		for _, v := range indexList {
			indexModelList = append(indexModelList, mongo.IndexModel{Keys: bson.D{{Key: v, Value: 1}}, Options: nil})

		}

		db.CreateCollection(ctx, colName)
		col := db.Collection(colName)
		indexes, err := col.Indexes().CreateMany(ctx, indexModelList)
		if err != nil {
			LogFatal("Mongo", "'%s' Index Create Fail: %s", colName, err.Error())
		}
		LogInfo("Mongo", "'%s' Index Create Success: %v", colName, indexes)
	}
}

func GetMongoCollection(db *mongo.Database, v interface{}) *mongo.Collection {
	t := reflect.TypeOf(v)
	p := pluralize.NewClient()
	plural := p.Plural(t.Name())
	name := strings.ToLower(plural[:1] + plural[1:])
	return db.Collection(name)
}

// 初始化数据库
func InitMongoDBConnect(mongoose *MongoConfig) (db *mongo.Database) {
	var url string
	if mongoose.Username != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mongoose.Username, mongoose.Password, mongoose.Url, mongoose.Port, mongoose.Database)
		if mongoose.Username == "admin" {
			url = fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoose.Username, mongoose.Password, mongoose.Url, mongoose.Port)
		}
	} else {
		url = fmt.Sprintf("mongodb://%s:%s/%s", mongoose.Url, mongoose.Port, mongoose.Database)
	}

	if strings.Contains(mongoose.Url, "mongodb://") {
		// 支持仅填写url
		url = mongoose.Url
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		LogFatal("Mongo", "%s %s", mongoose.Name, err.Error())
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		LogFatal("Mongo", "%s %s", mongoose.Name, err.Error())
	}

	db = client.Database(mongoose.Database)
	LogInfo("Mongo", "%s 连接成功", mongoose.Name)
	return
}

type MongoConfig struct {
	Name     string // 自定义名称
	Username string // 用户名
	Password string // 密码
	Url      string // url链接
	Port     string // 端口
	Database string // 数据库名称
}
