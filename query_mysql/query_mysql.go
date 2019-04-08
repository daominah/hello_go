package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	"math/rand"
)

const (
	MYSQL_ER_DUP_KEY = 1022
)

type Merchant struct {
	Id   int64
	Name string
}

type Voucher struct {
	Id       int64
	Name     string
	Merchant *Merchant `orm:"rel(fk)"`
}

type VouTicket struct {
	Id      int64
	Serial  string
	Voucher *Voucher `orm:"rel(fk)"`
}

var Ormer orm.Ormer

func init() {
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		"tungdt", "123qwe", "localhost", "3306", "minah_test")
	orm.RegisterDataBase("default", "mysql", dataSource)

	orm.RegisterModel(new(Merchant))
	orm.RegisterModel(new(Voucher))
	orm.RegisterModel(new(VouTicket))
	orm.RunSyncdb("default", false, true)

	Ormer = orm.NewOrm()
	_, err := Ormer.Raw(`
        ALTER TABLE voucher ADD CONSTRAINT fk_voucher_merchant
        FOREIGN KEY (merchant_id) REFERENCES merchant(id); `).Exec()
	beego.SetLogFuncCall(true)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		if mysqlErr.Number != MYSQL_ER_DUP_KEY {
			logs.Error(err)
		}
	}
	_, err = Ormer.Raw(`
        ALTER TABLE vou_ticket ADD CONSTRAINT fk_ticket_voucher
        FOREIGN KEY (voucher_id) REFERENCES voucher(id); `).Exec()
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		if mysqlErr.Number != MYSQL_ER_DUP_KEY {
			logs.Error(err)
		}
	}
}

func InsertTestData() {
	// insert merchants
	merchants := make([]*Merchant, 0)
	NMerchants := int64(5)
	for i := int64(1); i <= NMerchants; i++ {
		mer := &Merchant{Id: i, Name: fmt.Sprintf("Mer%v", i)}
		merchants = append(merchants, mer)
	}
	_, err := Ormer.InsertMulti(int(NMerchants), merchants)
	if err != nil {
		logs.Error(err)
	} else {
		logs.Info(fmt.Sprintf("Inserted %v merchants. \n", NMerchants))
	}
	// insert vouchers
	NVouBulks := int64(10)
	NVousPerBulk := int64(10)
	for j := int64(1); j <= NVouBulks; j++ {
		vouchers := make([]*Voucher, 0)
		for i := int64(1); i <= NVousPerBulk; i++ {
			id := NVousPerBulk*(j-1) + i
			vou := &Voucher{Id: id, Name: fmt.Sprintf("Vou%v", id),
				Merchant: &Merchant{Id: 1 + rand.Int63n(NMerchants)}}
			vouchers = append(vouchers, vou)
		}
		_, err = Ormer.InsertMulti(int(NVousPerBulk), vouchers)
		if err != nil {
			logs.Error(err)
		} else {
			logs.Info(fmt.Sprintf("%v Inserted %v vouchers. \n",
				j, NVousPerBulk))
		}
	}
	// insert vouTickets
	NTicketBulks := int64(2)
	NTicketsPerBulk := int64(10000)
	for j := int64(1); j <= NTicketBulks; j++ {
		vouTickets := make([]*VouTicket, 0)
		for i := int64(1); i <= NTicketsPerBulk; i++ {
			id := NTicketsPerBulk*(j-1) + i
			ticket := &VouTicket{Id: id, Serial: fmt.Sprintf("Tic%v", id),
				Voucher: &Voucher{Id: 1 + rand.Int63n(NVousPerBulk*NVouBulks)}}
			vouTickets = append(vouTickets, ticket)
		}
		_, err = Ormer.InsertMulti(int(NVousPerBulk), vouTickets)
		if err != nil {
			logs.Error(err)
		} else {
			logs.Info(fmt.Sprintf("%v Inserted %v tickets. \n",
				j, NVousPerBulk))
		}
	}
}

func main() {
	InsertTestData()
}
