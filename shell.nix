let
  sources = import ./nix/sources.nix;
in
{ pkgs ? import sources.nixpkgs { } }:
let
  sqlite-utils = pkgs.python3Packages.buildPythonPackage rec {
    name = "sqlite-utils";
    version = "git";
    src = sources.sqlite-utils;
    buildInputs = with pkgs.python3Packages; [ pytest black pytestrunner ];
    propagatedBuildInputs = with pkgs.python3Packages; [ click click-default-group tabulate ];
  };
in
pkgs.mkShell {
  name = "lcbo-car-fuel";
  buildInputs = [
    pkgs.curl
    pkgs.jq
    pkgs.fish
    pkgs.pv
    pkgs.sqlite
    pkgs.python3
    sqlite-utils
    pkgs.metabase
  ];
}
