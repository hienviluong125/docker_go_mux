SELECT 'CREATE DATABASE docker_go_mux'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'docker_go_mux')\gexec
