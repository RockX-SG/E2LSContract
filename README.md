### generate golang code
````
abigen --abi ./stake.abi --pkg staking_contract --out ./stake_contract/staking_contract.go
````

### get abi
```
cat RockXStaking.json | jq .abi
```