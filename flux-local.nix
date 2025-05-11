{
  buildPythonPackage,
  fetchPypi,

  # propagatedBuildInputs
  pyyaml,
  gitpython,
  aiofiles,
  mashumaro,
  nest-asyncio,
  pytest,
  pytest-asyncio,
}:
buildPythonPackage rec {
  pname = "flux_local";
  version = "7.5.0";

  src = fetchPypi {
    inherit pname version;
    hash = "sha256-fWaZmim6NNGptiiDWkNnvJ5b3Lz0hlojXc25Gc8hYHY=";
  };

  propagatedBuildInputs = [
    pyyaml
    gitpython
    aiofiles
    mashumaro
    nest-asyncio
    pytest
    pytest-asyncio
  ];
}
