include:
  - hwaas-site

supervisor:
  pkg.installed
    - require:
      - sls: hwwas-site
