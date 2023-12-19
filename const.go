package cmcpro

import "time"

const ApiPoint = "https://pro-api.coinmarketcap.com"
const ApiPointTest = "https://sandbox-api.coinmarketcap.com"

const All = "all"
const Coins = "coins"
const Tokens = "tokens"
const Fees = "fees"
const NoFees = "no_fees"

const SortAsc = "asc"
const SortDesc = "desc"

const Timeout = time.Second * 15

const CryptocurrencyMaxLimit = 5000
const ExchangeMaxLimit = 5000
