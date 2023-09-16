To generate a network with the second choice option, you need to use the `config` parameter.
E.g.
```bash
cd generate_network_data
go build
./generate_network_data -random=false -config=true -conffile=../config.yaml
```