'use strict';

var util       = require('util'),
    path       = require('path'),
    generators = require('yeoman-generator'),
    _          = require('lodash'),
    _s         = require('underscore.string'),
    pluralize  = require('pluralize'),
    asciify    = require('asciify'),
    mkdir      = require('fs').mkdirSync,
    exists     = require('fs').existsSync;


function mkdirIfNotExists(path){
  if (! exists(path)){
    mkdir(path)
  }
}

module.exports = generators.Base.extend({
  constructor: function () {
    generators.Base.apply(this, arguments);

    this.option('flat', {
      type: Boolean,
      required: false,
      defaults: false,
      desc: 'When specified, generators will be created at the top level of the project.'
    });
  },

  initializing: function () {
    this.config.set('structure', this.options.flat ? 'flat' : 'nested');
    this.generatorsPrefix = this.options.flat ? '' : 'generators/';
    this.appGeneratorDir = this.options.flat ? 'app' : 'generators';
  },

  prompting: {
    askFor: function () {
      var done = this.async();
      this.log('\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '|g|o| |r|e|s|t| |g|e|n|e|r|a|t|o|r|\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '\n'
      );

      var prompts = [{
        type: 'input',
        name: 'baseName',
        message: 'What is the name of your application?',
        store   : true,
        default: 'myapp'
      },
        {
          type: 'input',
          name: 'basePackage',
          message: 'What is the name of your base package mine is github.com/maleck13?\n your app will live in $GOPATH/src/<basePackage>/<appName>',
          store   : true,
          required: true
        },
      {
        type: 'input',
        name: 'prometheus',
        message: 'Do you want to include prometheus metrics? [Yes]',
        store   : true,
        default: "yes"
      },
        {
          type: 'input',
          name: 'stomp',
          message: 'Do you want to include an stomp messaging boilerplate (yes/no). [no]',
          store   : true,
          default: "no"
        },
        {
          type: 'input',
          name: 'database',
          message: 'Do you want to include a database boilerplate. (mongo/no) [no]',
          store   : true,
          default: "no"
        }];

      this.prompt(prompts, function (props) {
        this.baseName = props.baseName;
        this.basePackage = props.basePackage;
        this.metrics = props.prometheus;
        this.database = props.database;
        this.stomp = props.stomp;

        done();
      }.bind(this));
    }
  },

  writing: {
    app: function () {
      var database = this.database;
      var cmdDir = "cmd/";
      var configDir = "config/";
      var apiDir = "api/";
      var apiMW = "api/middleware";
      var domainDir = "domain/";

      this.copy('gitignore', '.gitignore');

      mkdirIfNotExists(cmdDir);
      mkdirIfNotExists(configDir);
      mkdirIfNotExists(apiDir);
      mkdirIfNotExists(apiMW);
      mkdirIfNotExists(domainDir);

      var commonTemplateVals = {
        basePackage: this.basePackage,
        baseName: this.baseName,
        database: this.database,
        metrics: this.metrics,
        messaging: this.stomp
      };
      this.fs.copyTpl(
        this.templatePath('_main.go'),
        this.destinationPath('main.go'),
        commonTemplateVals
      );
      this.fs.copyTpl(
        this.templatePath('cmd/_server.go'),
        this.destinationPath('cmd/server.go'),
        commonTemplateVals
      );

      this.fs.copyTpl(
        this.templatePath('config/_config.go'),
        this.destinationPath('config/config.go'),
        commonTemplateVals
      );
      this.fs.copyTpl(
        this.templatePath('config/config.json'),
        this.destinationPath('config/config.json'),
        commonTemplateVals
      );

      if("no" !== database){
        mkdirIfNotExists("data/");
        mkdirIfNotExists("data/model");
      }

      if("mongo" == database){
        this.fs.copyTpl(
          this.templatePath('data/_database_mongo.go'),
          this.destinationPath('data/mongo.go'),
          commonTemplateVals
        );

        this.fs.copyTpl(
          this.templatePath('data/model/_example.go'),
          this.destinationPath('data/model/example.go'),
          commonTemplateVals
        );

        this.fs.copyTpl(
          this.templatePath('data/model/_example.go'),
          this.destinationPath('data/model/example.go'),
          commonTemplateVals
        );

        this.fs.copyTpl(
          this.templatePath('domain/_exampleService.go'),
          this.destinationPath('domain/exampleService.go'),
          commonTemplateVals
        );

        this.fs.copyTpl(
          this.templatePath('domain/_exampleService_test.go'),
          this.destinationPath('domain/exampleService_test.go'),
          commonTemplateVals
        );

      }

      if("yes" == this.stomp){
        this.fs.copyTpl(
          this.templatePath('data/_stomp.go'),
          this.destinationPath('data/stomp.go'),
          commonTemplateVals
        );
      }

      //api
      this.fs.copyTpl(
        this.templatePath('api/_indexHandler.go'),
        this.destinationPath('api/indexHandler.go'),
        commonTemplateVals
      );
      this.fs.copyTpl(
        this.templatePath('api/_router.go'),
        this.destinationPath('api/router.go'),
        commonTemplateVals
      );
      this.fs.copyTpl(
        this.templatePath('api/_routeErrorHandler.go'),
        this.destinationPath('api/routeErrorHandler.go'),
        commonTemplateVals
      );

      this.fs.copyTpl(
        this.templatePath('api/_indexHandler_test.go'),
        this.destinationPath('api/indexHandler_test.go'),
        commonTemplateVals
      );


      //docker
      this.fs.copyTpl(
        this.templatePath('Dockerfile'),
        this.destinationPath('Dockerfile'),
        commonTemplateVals
      );

      this.fs.copyTpl(
        this.templatePath('docker-compose.yml'),
        this.destinationPath('docker-compose.yml'),
        commonTemplateVals
      );

      this.template('domain/Readme.md', 'domain/Readme.md');
      this.template('api/_errors.go', 'api/errors.go');
      this.template('api/_healthHandler.go', 'api/healthHandler.go');
      this.template('api/middleware/_example.go','api/middleware/example.go');
      this.template('README.md', 'README.md');
    }
  }
});
