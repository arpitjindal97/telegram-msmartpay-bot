set :docker_repo, ENV["DOCKER_IMG"]

set :server_host, ENV["server_1_host"]
set :server_user, ENV["server_1_user"]
set :server_url, "#{fetch(:server_user)}@#{fetch(:server_host)}"

task :deploy do
	on "#{fetch(:server_url)}" do
		execute "export docker_cont_id=$(docker ps -a | grep #{fetch(:docker_repo)} | cut -d ' ' -f 1 ) 
		if [ \"$docker_cont_id\" == \"\" ];then exit 0
		else 	docker stop $docker_cont_id  
			docker rm $docker_cont_id
			docker rmi #{fetch(:docker_repo)}
		fi"
		execute "docker pull #{fetch(:docker_repo)} > /dev/null"
		execute "docker run -d #{fetch(:docker_repo)}"
	end
end
