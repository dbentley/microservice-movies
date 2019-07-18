for f in listdir('k8s'):
  k8s_yaml(f)


docker_build('users-db', 'services/users/src/db')
docker_build('movies-db', 'services/movies/src/db')

docker_build('users', 'services/users',
             live_update=[
               sync('services/users', '/usr/src/app'),
               ])
docker_build('movies', 'services/movies',
             live_update=[
               sync('services/movies', '/usr/src/app'),
               ])
docker_build('frontend', 'services/web',
             live_update=[
               sync('services/web', '/usr/src/app'),
               ])
docker_build('ratings', 'services/ratings',
             live_update=[
               sync('services/ratings', '/go/src/github.com/dbentley/microservice-movies/services/ratings'),
               run('go install github.com/dbentley/microservice-movies/services/ratings'),
               restart_container(),
               ])

k8s_resource('users', port_forwards=3000)
k8s_resource('movies', port_forwards=3001)
k8s_resource('ratings', port_forwards=3002)
k8s_resource('frontend', port_forwards=3006)
