{
  description = "LCBO Car Fuel Index";
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };
  outputs = { self, ... }@inputs:
    (inputs.flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import inputs.nixpkgs { inherit system; };
        in
        {
          devShells.default = pkgs.mkShell rec {
            name = "lcbo-car-fuel";
            buildInputs = with pkgs; [
              gnumake
              curl
              jq
              fish
              pv
              moreutils
              go_1_18
              sqlite
              sqlite-utils
              (python3.withPackages (ps: with ps; [
                ps.pandas
                ps.pyarrow
                ps.fastparquet
                ps.python-snappy
                ps.ipython
                ps.requests
                ps.backoff
              ]))
              poetry
            ];
          };
        }));
}
