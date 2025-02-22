package cruds

import (
	"bondi-push-notification/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func Initialize() error {
	orm.RegisterModel(new(models.PushSubscribers))
	orm.RegisterModel(new(models.StudentCourse))
	orm.RegisterModel(new(models.PushNotification))

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		return err
	}

	user, _ := beego.AppConfig.String("PG::User")
	password, _ := beego.AppConfig.String("PG::Password")
	dbName, _ := beego.AppConfig.String("PG::DbName")
	host, _ := beego.AppConfig.String("PG::Host")
	port, _ := beego.AppConfig.String("PG::Port")
	dns := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=require", user, password, dbName, host, port)
	// dns := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbName, host, port)
	if err := orm.RegisterDataBase("default", "postgres", dns); err != nil {
		return err
	}
	if err := orm.RunSyncdb("default", false, true); err != nil {
		return err
	}
	return nil
}
