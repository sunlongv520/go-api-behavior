package db

import "fmt"

type Order struct {
	Order_id int `db:"order_id"`
}

type OrderItems struct {
	GoodsName string `db:"goods_name"`
}

type OrderList struct {
	Order
	OrderItems
	BrandName string `db:"brand_name"`
}

func Findabc(order_id int)  (data *OrderList,err error){
	data  = &OrderList{}
//	sql := `select a.order_id,b.goods_name,b.brand_name from lie_order as a left join
//lie_order_items as b  on a.order_id = b.order_id
//where a.create_time >= 1550115738`
	sql := `select a.order_id,b.goods_name,b.brand_name from lie_order as a left join
lie_order_items as b  on a.order_id = b.order_id
where a.order_id = ?`
	err = DB.Get(data,sql,order_id)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("...............")
	fmt.Printf("%+v",data)
	fmt.Println("...............")
	//for i ,v := range data{
	//	fmt.Println(i)
	//	fmt.Println(v)
	//}
	//DB.Close()
	return
}
