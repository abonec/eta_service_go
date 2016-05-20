# Modes

Eta service has four modes selected by `-mode=xxx` argument

* eta_server [default]
* update_cab_server
* migrate
* send_message [deprecated]

# eta_server
    ./eta_service_go

# update_cab_server
    ./eta_service_go -mode=update_cab_server

# migrate
    ./eta_service_go -mode=migrate

# send_message
    ./eta_service_go -mode=send_message -message={...}
