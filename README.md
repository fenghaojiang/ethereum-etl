## Ethereum-ETL

Ethereum-ETL is a ETL tool to sync data with EVM-like block chain.  
You can configure the destination endpoints like `MySQL` / `Postgresql` / `MongoDB` / `TiDB` ...  

### 
[![Test](https://github.com/fenghaojiang/ethereum-etl/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/fenghaojiang/ethereum-etl/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/fenghaojiang/ethereum-etl?status.png)](https://godoc.org/github.com/fenghaojiang/ethereum-etl)
[![GoReportCard](https://goreportcard.com/badge/github.com/fenghaojiang/ethereum-etl)](https://goreportcard.com/report/github.com/fenghaojiang/ethereum-etl)


### Goals  

This project aims to solve the pain points and difficulties encountered when pulling data from EVM-based public chains such as Ethereum for ETL purposes, helping users synchronize on-chain data and supporting data synchronization to various types of databases such as OLAP/OLTP/HTAP.  

#### Key Features 

- Resolving [reorg issue](https://www.alchemy.com/overviews/what-is-a-reorg) - when reorg occurs, the data already consumed by the database will be corrected/modified. 

![reorg-handle](./images/reorg_handle.png)  

- Real-time updates - users can obtain the latest data.
- Provide blazing-fast solution to synchronize data.
- Support for multiple databases to meet users' various scenario requirements.