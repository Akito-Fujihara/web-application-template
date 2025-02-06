module.exports = {
  public: {
    input: {
      target: './doc/api/openapi/openapi.publicapi.yaml',
    },
    output: {
      mode: 'split',
      target: './frontend/src/hooks/public/clients',
      schemas: './frontend/src/hooks/public/models',
      client: 'swr',
      override: {
        mutator: {
          path: './frontend/src/hooks/public/publicInstance.ts',
          name: 'publicInstance',
        },
      },
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
  private: {
    input: {
      target: './doc/api/openapi/openapi.privateapi.yaml',
    },
    output: {
      mode: 'split',
      target: './frontend/src/hooks/private/clients',
      schemas: './frontend/src/hooks/private/models',
      client: 'axios',
      override: {
        mutator: {
          path: './frontend/src/hooks/private/privateInstance.ts',
          name: 'privateInstance',
        },
      },
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
}
