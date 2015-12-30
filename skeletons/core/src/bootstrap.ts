import {bootstrap} from 'angular2/platform/browser';
import {App} from './app/app';

function main() {
  return bootstrap(App);
}

document.addEventListener('DOMContentLoaded', main);
