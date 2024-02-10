build:
	sam build

local_invoke:
	sam local invoke --env-vars env.json
