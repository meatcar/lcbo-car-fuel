#!/usr/bin/env python3
import sys
import os
import pandas as pd
import requests
import backoff

DB_PATH = os.environ['DB_PATH']
API_BASE = os.environ['API_BASE']

@backoff.on_exception(backoff.expo,
                      (requests.exceptions.Timeout,
                       requests.exceptions.ConnectionError),
                      max_time=60)
def get_url(url):
    print(url, file=sys.stderr)
    return requests.get(url)

def read_parquet(name):
    return pd.read_parquet('{}/{}.parquet'.format(DB_PATH, name))

def get_all_store_stock(name):
    df = read_parquet('stores')
    stores = df['locationNumber']

    df = read_parquet('products')
    skus = df[df['itemName'].str.startswith(name)]['itemNumber']
    sku_list = ','.join([str(x) for x in skus])
    if sku_list == '':
        raise Exception('"{}" does not match any items'.format(name))

    stock = pd.DataFrame()
    for n in stores:
        url = '{}/v7/products/?locationNumber={}&skuList={}'.format(API_BASE, n, sku_list)
        resp = get_url(url)

        if resp.status_code != 200:
            raise Exception('API Error: {}'.format(resp.text))

        data = resp.json()
        df = pd.DataFrame(data['products'])
        df['itemNumber'] = pd.to_numeric(df['itemNumber'])
        df['locationNumber'] = n
        stock = stock.append(df[['productQuantity', 'itemNumber', 'locationNumber']])

    return stock

if __name__ == '__main__':
    name = sys.argv[1]
    stock_file = '{}/{}.parquet'.format(DB_PATH, name)
    stock = get_all_store_stock(name)
    stock.to_parquet(stock_file)
    print('done', file=sys.stderr)
