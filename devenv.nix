{
  pkgs,
  lib,
  ...
}: let
  util = import ./nix/util.nix {inherit pkgs lib;};
  inherit (util) enableAll;
in {
  packages = with pkgs; [spec-kit go-task protobuf grpcurl buf protoc-gen-go protoc-gen-connect-go];
  languages.go = {
    enable = true;
    version = "1.25.5";
  };
  languages.javascript = {
    enable = true;
    package = pkgs.nodejs_24;
    npm.enable = true;
  };

  # https://devenv.sh/processes/
  # processes.dev.exec = "${lib.getExe pkgs.watchexec} -n -- ls -la";

  # https://devenv.sh/services/
  # services.postgres.enable = true;

  scripts.startup.exec = ''
    echo  "Golang + Angular + NX Monorepo Development Environment"
    echo  "NodeJS $(node --version)"
    echo  "NPM $(npm --version)"
    git   --version
    go    version
  '';

  # https://devenv.sh/basics/
  enterShell = ''
    startup
  '';

  # https://devenv.sh/tasks/
  # tasks = {
  #   "myproj:setup".exec = "mytool build";
  #   "devenv:enterShell".after = [ "myproj:setup" ];
  # };

  # https://devenv.sh/tests/
  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
  '';

  # https://devenv.sh/git-hooks/
  git-hooks = {
    hooks = enableAll ["shellcheck" "action-validator" "actionlint" "commitizen"] {
      shellcheck.excludes = ["^\\.specify/"];
    };
  };
}
