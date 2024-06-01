go install github.com/spf13/cobra-cli@latest

cd /workspaces/home-config

task build

echo "export PATH=\$PATH:/workspaces/home-config/bin" >> ~/.bashrc

echo "source <(home-cli completion bash)" >> ~/.bashrc