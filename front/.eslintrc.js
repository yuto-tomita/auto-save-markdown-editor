module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    "plugin:vue/vue3-essential",
    "plugin:vue/vue3-recommended",
    "eslint:recommended",
    "@vue/typescript/recommended",
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  rules: {
		"semi": "error",
    "vue/html-indent": "error",
    "vue/html-quotes": "error",
    "vue/html-end-tags": "error",
    "vue/html-self-closing": "error",
  }
};