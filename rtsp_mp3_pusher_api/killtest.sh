kill -9 `ps -ef | grep "./PusherModuleTest" | awk 'NR==2{print $2}' `
