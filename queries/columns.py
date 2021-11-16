#!/usr/bin/env python3
import pandas as pd
import sys

df = pd.read_parquet(sys.argv[1])
print(df.columns)
