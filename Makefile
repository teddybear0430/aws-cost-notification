build:
	sam build

local_invoke:
	sam local invoke --parameter-overrides LineAccessToken=$LINE_ACCESS_TOKEN OpenExchangeRatesAppId=$OPEN_EXCHANGE_RATES_APP_ID

deploy:
	sam deploy --parameter-overrides LineAccessToken=$LINE_ACCESS_TOKEN OpenExchangeRatesAppId=$OPEN_EXCHANGE_RATES_APP_ID
