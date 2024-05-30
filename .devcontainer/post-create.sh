go install github.com/spf13/cobra-cli@latest

cd /workspaces/home-config

task build

echo "export PATH=\$PATH:/workspaces/home-config/bin" >> ~/.bashrc

echo "source <(home-cli completion bash)" >> ~/.bashrc

docker context create athena --docker "host=ssh://mgarnier@100.64.98.100"
docker context create apollon --docker "host=ssh://mgarnier@100.64.98.99"
docker context create zephyr --docker "host=ssh://mgarnier@100.64.98.97"
