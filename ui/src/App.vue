<template>
  <div id="app">
    <b-jumbotron header="ServerCheck" lead="Check SSL domain status, servers and more!" />
    <b-nav tabs>
      <b-nav-item @click="changeView('search')" name="search" :active="isSearch()">
        Search domain
      </b-nav-item>
      <b-nav-item @click="changeView('list')" name="list" :active="isList()">
        List searched domains
      </b-nav-item>
    </b-nav>
    <SearchDomain v-if="isSearch()" />
    <ListSearched v-if="isList()" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import SearchDomain from './components/SearchDomain.vue';
import ListSearched from './components/ListSearched.vue';

type IView = 'search' | 'list';

@Component({
  components: {
    SearchDomain,
    ListSearched,
  },
})
export default class App extends Vue {
  public view: IView = 'search';

  public changeView(view: IView) {
    this.view = view;
  }

  public isSearch(): boolean {
    return this.view === 'search';
  }

  public isList(): boolean {
    return this.view === 'list';
  }
}
</script>

<style>
  #app .jumbotron {
    padding: 1rem;
  }
</style>
