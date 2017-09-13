set :docker_repo, "$DOCKER_USERNAME/telegram-msmartpay"
set :server, "$server_1_host"
set :server_user, "$server_1_user"
server "#{fetch(:server)}", user: "#{fetch(:server_user)}"
task :deploy do
	on roles(:all) do |server|
		execute "export docker_cont_id=$(docker ps -a | grep #{fetch(:docker_repo)} | cut -d ' ' -f 1 ) 
		if [ \"$docker_cont_id=\" == \"\" ];then exit 0
		else 	docker stop $docker_cont_id  
			docker rm $docker_cont_id
			docker rmi #{fetch(:docker_repo)}
		fi"
		execute "docker pull #{fetch(:docker_repo)} > /dev/null"
		execute "docker run -d #{fetch(:docker_repo)}"
	end
end
