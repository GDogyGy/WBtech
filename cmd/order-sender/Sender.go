package main

import (
	"github.com/nats-io/stan.go"
	"log/slog"
	"os"
)

func main() {
	data := []string{`{
    "order_uid": "b563feb7b2b84b6test1",
    "track_number": "WBILMTESTTRACK1",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov1",
      "phone": "+97200000001",
      "zip": "26398091",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test1@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test1",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934930,
        "track_number": "WBILMTESTTRACK1",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      },
     {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK2",
        "price": 457,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 357,
        "nm_id": 23892121,
        "brand": "Adidas",
        "status": 212
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`, `{
    "order_uid": "b563feb7b2b84b6test2",
    "track_number": "WBILMTESTTRACK2",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov2",
      "phone": "+97200000002",
      "zip": "2639809",
      "city": "Moscow",
      "address": "Lenina 16",
      "region": "Kraiot",
      "email": "test2@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1815,
      "payment_dt": 16379077272,
      "bank": "alpha",
      "delivery_cost": 1502,
      "goods_total": 319,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934930,
        "track_number": "WBILMTESTTRACK2",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      },
      {
        "chrt_id": 9934931,
        "track_number": "WBILMTESTTRACK2",
        "price": 457,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 357,
        "nm_id": 23892121,
        "brand": "Adidas",
        "status": 212
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }`,
		`{
		"order_uid": "b563feb7b2b84b6test3",
		"track_number": "WBILMTESTTRACK3",
		"entry": "WBIL",
		"delivery": {
			"name": "Test Testov3",
			"phone": "+97200000003",
			"zip": "2639809",
			"city": "St. Peter",
			"address": "Ploshad Mira 15",
			"region": "Kraiot",
			"email": "test3@gmail.com"
		},
		"payment": {
			"transaction": "b563feb7b2b84b6test3",
			"request_id": "",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 1927,
			"payment_dt": 16379077273,
			"bank": "alpha",
			"delivery_cost": 4500,
			"goods_total": 217,
			"custom_fee": 0
		},
		"items": [
	{
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK3",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	}
	],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`,
		`{
		"order_uid": "b563feb7b2b84b6test4",
		"track_number": "WBILMTESTTRACK4",
		"entry": "WBIL",
		"delivery": {
			"name": "Ivan ivanov",
			"phone": "+97200000004",
			"zip": "2639809",
			"city": "Vladimir",
			"address": "Komsomolec 15",
			"region": "Lenina",
			"email": "test4@gmail.com"
		},
		"payment": {
			"transaction": "b563feb7b2b84b6test4",
			"request_id": "",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 3817,
			"payment_dt": 16379077274,
			"bank": "alpha",
			"delivery_cost": 500,
			"goods_total": 517,
			"custom_fee": 0
		},
		"items": [
	{
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK4",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 217,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	}
	],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2022-11-26T06:22:19Z",
		"oof_shard": "1"
	}`,
		`{
		"order_uid": "b563feb7b2b84b6test5",
		"track_number": "WBILMTESTTRACK5",
		"entry": "WBIL",
		"delivery": {
			"name": "Pavel Ivanov",
			"phone": "+9720000000",
			"zip": "2639809",
			"city": "Ivanovo",
			"address": "Lenina 2",
			"region": "Lininsiy",
			"email": "test5@gmail.com"
		},
		"payment": {
			"transaction": "b563feb7b2b84b6test5",
			"request_id": "",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 1817,
			"payment_dt": 1637907727,
			"bank": "alpha",
			"delivery_cost": 1500,
			"goods_total": 317,
			"custom_fee": 0
		},
		"items": [
	{
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK5",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	}
	],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`,
	}

	var log *slog.Logger
	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	sc, err := stan.Connect("test_cluster", "stan-server1")
	if err != nil {
		log.Error("Failed nats connect", slog.StringValue(err.Error()))
	}
	defer sc.Close()

	for _, v := range data {
		if err := sc.Publish("updates", []byte(v)); err != nil {
			log.Error("Failed mess mailer ", slog.StringValue(err.Error()))
		}
	}
}
