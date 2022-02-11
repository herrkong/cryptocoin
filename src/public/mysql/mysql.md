


recoverkline

datetime比较

SELECT * FROM options_trade WHERE instrument_id LIKE 'BTCUSDT-20211231-43000-C' AND trade_time > '2021-12-23 11:53:00.931'



# left join otfr
/* SELECT 
    otf.* 
    otfr.operate_time
    otfr.operator
FROM 
    options_trading_fee otf
LEFT JOIN 
    options_trading_fee_record otfr
ON
    otf.id = otfr.ref_id

# left join ou
SELECT
	otf.*,
	ou.order_max_fee,
	ou.insurance_rate,
	ou.settle_fee,
	ou.settle_buy_max_fee,
	ou.settle_sell_max_fee
FROM
	options_trading_fee otf
LEFT JOIN 
    options_underlying ou 
ON 
    otf.underlying_id = ou.underlying_id */


# 多表left join
SELECT 
    otf.*,
    ou.order_max_fee,
	ou.insurance_rate,
	ou.settle_fee,
	ou.settle_buy_max_fee,
	ou.settle_sell_max_fee,
    otfr.operate_time,
    otfr.operator
FROM 
    options_trading_fee otf
    LEFT JOIN options_underlying ou ON otf.underlying_id = ou.underlying_id
    LEFT JOIN options_trading_fee_record otfr ON otf.id = otfr.ref_id





// options_trade表中的交易数据 （需要做分库分表?） 之后trade可能会非常大

select instrument_id,count(*) from options_trade group by instrument_id;




#### 创建数据库

