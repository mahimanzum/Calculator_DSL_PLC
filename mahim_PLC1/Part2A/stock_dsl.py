import json
from pprint import pprint
import sys
def read_json(path):
    f = open(path)
    data = json.load(f)
    #pprint(data)
    return data 
def write_file(content, path):
    f = open(path, "w")
    f.write(content)
    f.close()

#data = read_json("stocks.json")
file_name = sys.argv[1]
data = read_json(file_name)
#print(data.keys())
#INSERT INTO BuyRequests (NumShares, Symbol, MaxPrice, AccountID) VALUES (‘100’, ‘IBM’, ‘45’,  ‘Hokie123’)

#for part 2A
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
write_file(make_dsl(),file_name.replace("json", "dsl"))
write_file(make_sql(),file_name.replace("json", "sql"))







#for part 2B
'''
<stock_trade_requests> →  [delete]? ‘(' <trade> {‘,’ <trade>} ‘) for account' <acct_ident>’.’
<trade> →  <number> <stock_symbol> ‘shares’ (‘buy at max' | ‘sell at min') <number>
<number> →  [1-9] {[0-9]}
<stock_symbol> →
 'AAPL'|'HP'|'IBM'|'AMZN'|'MSFT'|'GOOGL'|'INTC'|'CSCO'|'ORCL'|'QCOM'
<acct_ident> →  ‘“‘alpha_char { alpha_char | digit | ’_’} ‘“‘

Note:  ‘“‘ is a “ surrounded by ‘

'''


def make_sql_with_delete():
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
    if 'delete' in data.keys():
        if 'buy' in data['delete'].keys():
            for item in data['delete']['buy']:
                if 'at max' in item.keys():
                    sql_file.append("DELETE FROM BuyRequests WHERE (NumShares = '{}' AND Symbol = '{}' AND MaxPrice = '{}' AND AccountID = '{}')".format(item['shares'], item['stock symbol'], item['at max'], data['user id']))
                elif 'at min' in item.keys():
                    sql_file.append("DELETE FROM BuyRequests WHERE (NumShares = '{}' AND Symbol = '{}' AND MaxPrice = '{}' AND AccountID = '{}')".format(item['shares'], item['stock symbol'], item['at min'], data['user id']))
        if 'sell' in data['delete'].keys():
            for item in data['delete']['sell']:
                if 'at max' in item.keys():
                    sql_file.append("DELETE FROM SellRequests WHERE (NumShares = '{}' AND Symbol = '{}' AND MaxPrice = '{}' AND AccountID = '{}')".format(item['shares'], item['stock symbol'], item['at max'], data['user id']))
                elif 'at min' in item.keys():
                    sql_file.append("DELETE FROM SellRequests WHERE (NumShares = '{}' AND Symbol = '{}' AND MaxPrice = '{}' AND AccountID = '{}')".format(item['shares'], item['stock symbol'], item['at min'], data['user id']))
    return "\n".join(sql_file)

def make_dsl_with_delete():
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
    delete_trades = []       
    if 'delete' in data.keys():
        if 'buy' in data['delete'].keys():
            for item in data['delete']['buy']:
                if 'at max' in item.keys():
                    trade = "{} {} shares buy at max {}".format(item['shares'],item['stock symbol'], item['at max'])
                    delete_trades.append(trade)
                elif 'at min' in item.keys():
                    trade = "{} {} shares buy at min {}".format(item['shares'],item['stock symbol'], item['at min'])
                    delete_trades.append(trade)
        if 'sell' in data['delete'].keys():
            for item in data['delete']['sell']:
                if 'at max' in item.keys():
                    trade = "{} {} shares sell at max {}".format(item['shares'],item['stock symbol'], item['at max'])
                    delete_trades.append(trade)
                elif 'at min' in item.keys():
                    trade = "{} {} shares sell at min {}".format(item['shares'],item['stock symbol'], item['at min'])
                    delete_trades.append(trade)
    
    
    stock_trade_requests = "("+ ", ".join(trades)+")" +" for account {}".format(data['user id'])
    delete_stock_trade_requests = ""
    if len(delete_trades):
        delete_stock_trade_requests = "delete " + "("+ ", ".join(delete_trades)+")" +" for account {}".format(data['user id'])
    #print(stock_trade_requests)
    return stock_trade_requests+'\n'+delete_stock_trade_requests

#These functions are tested with delete command also    

#write_file(make_dsl_with_delete(),file_name.replace("json", "dsl"))
#write_file(make_sql_with_delete(),file_name.replace("json", "sql"))