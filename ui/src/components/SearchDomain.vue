<template>
  <div>
  <b-input-group>
    <b-form-input v-model="domain"></b-form-input>

    <b-input-group-append>
      <b-button variant="outline-secondary" @click="search()">Search</b-button>
    </b-input-group-append>
  </b-input-group>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class SearchDomain extends Vue {
  public domain = 'truora.com';
  public endpoint = '';
  public domainData: any;

  private getEndpoint() {
    return 'http://localhost:8005/domains/' + this.domain;
  }

  private mounted() {
    this.search();
  }

  private search() {
    fetch(this.getEndpoint())
      .then((response) => {
        return response.json();
      }).then((result) => {
        this.domainData = result;
      });
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
