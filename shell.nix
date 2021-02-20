let
  sources = import ./nix/sources.nix;
in
{ pkgs ? import sources.nixpkgs { } }:
pkgs.mkShell {
  name = "lcbo-car-fuel";
  buildInputs = [
    pkgs.niv
    pkgs.curl
    pkgs.jq
    pkgs.fish
    pkgs.pv
    pkgs.sqlite
    pkgs.python3
    pkgs.sqlite
    pkgs.sqlite-utils
    pkgs.metabase
  ];
}
