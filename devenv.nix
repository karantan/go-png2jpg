{ pkgs, ... }:

{
  packages = [
    pkgs.git
    pkgs.nodePackages.serverless
  ];

  languages.go.enable = true;
}
