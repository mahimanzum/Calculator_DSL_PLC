import json
from pprint import pprint
def read_json(path):
    f = open(path)
    data = json.load(f)
    #pprint(data)
    return data 
def write_file(content, path):
    f = open(path, "w")
    f.write(content)
    f.close()

data = read_json("stocks.json")
#print(data.keys())
#INSERT INTO BuyRequests (NumShares, Symbol, MaxPrice, AccountID) VALUES (‘100’, ‘IBM’, ‘45’,  ‘Hokie123’)
def make_sql():
    sql_file = []
    if 'buy' in data.keys():
        for item in data['buy']:
            if 'at max' in item.keys():
                sql_file.append("INSERT INTO BuyRequests (NumShares, Symbol, MaxPrice, AccountID) VALUES ('{}', '{}','{}', '{}')".format(item['shares'], item['stock symbol'], item['at max'], data['user id']))
            elif 'at min' in item.keys():
                sql_file.append("INSERT INTO BuyRequests (NumShares, Symbol, MinPrice, AccountID) VALUES ('{}', '{}','{}', '{}')".format(item['shares'], item['stock symbol'], item['at min'], data['user id']))
    if 'sell' in data.keys():
        for item in data['sell']:
            if 'at max' in item.keys():
                sql_file.append("INSERT INTO SellRequests (NumShares, Symbol, MaxPrice, AccountID) VALUES ('{}', '{}','{}', '{}')".format(item['shares'], item['stock symbol'], item['at max'], data['user id']))
            elif 'at min' in item.keys():
                sql_file.append("INSERT INTO SellRequests (NumShares, Symbol, MinPrice, AccountID) VALUES ('{}', '{}','{}', '{}')".format(item['shares'], item['stock symbol'], item['at min'], data['user id']))
    return "\n".join(sql_file)

def make_dsl():
    trades = []
    if 'buy' in data.keys():
        for item in data['buy']:
            if 'at max' in item.keys():
                trade = "{} {} shares buy at max {}".format(item['shares'],item['stock symbol'], item['at max'])
                trades.append(trade)
            elif 'at min' in item.keys():
                trade = "{} {} shares buy at min {}".format(item['shares'],item['stock symbol'], item['at min'])
                trades.append(trade)
    if 'sell' in data.keys():
        for item in data['sell']:
            if 'at max' in item.keys():
                trade = "{} {} shares sell at max {}".format(item['shares'],item['stock symbol'], item['at max'])
                trades.append(trade)
            elif 'at min' in item.keys():
                trade = "{} {} shares sell at min {}".format(item['shares'],item['stock symbol'], item['at min'])
                trades.append(trade)
    
    stock_trade_requests = "("+ ", ".join(trades)+")" +" for account {}".format(data['user id'])
    #print(stock_trade_requests)
    return stock_trade_requests
write_file(make_dsl(), "my_stocks.dsl")
write_file(make_sql(), "my_stocks.sql")
    

