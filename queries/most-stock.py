#!/usr/bin/env python3
import sys
import os
import pandas as pd
import requests

DB_PATH = os.environ['DB_PATH']
API_BASE = os.environ['API_BASE']

def read_parquet(name):
    return pd.read_parquet('{}/{}.parquet'.format(DB_PATH, name))

def get_most_stock(name):
    stores = read_parquet('stores')
    products = read_parquet('products')
    stock = read_parquet(name)

    stock['itemNumber'] = pd.to_numeric(stock['itemNumber'])
    stock['productQuantity'] = pd.to_numeric(stock['productQuantity'])
    stock = stock.sort_values('productQuantity', ascending=False)

    merged = pd.merge(stock, products, on='itemNumber', how='left')
    merged = pd.merge(merged, stores, on='locationNumber', how='left')
    return merged

if __name__ == '__main__':
    name = sys.argv[1]
    df = get_most_stock(name)
    df.to_excel('{}/{}.xlsx'.format(DB_PATH, name))
    print(df.head(20)[['productQuantity_x', 'itemName', 'locationName', 'locationCityName']])
    print('done', file=sys.stderr)

