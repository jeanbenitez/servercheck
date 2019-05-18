<template>
  <DomainTable :domains="domains" :loading="loading" />
</template>

<script lang="ts">
import { Component, Prop, Vue, Model } from 'vue-property-decorator';
import { Domain, DomainResponse } from '../interfaces/domain';
import { fetchJSON } from '../utils/http';
import DomainTable from './DomainTable.vue';

@Component({
  components: {
    DomainTable,
  },
})
export default class ListSearched extends Vue {
  private endpoint = 'http://localhost:8005/domains/';
  private domains: Domain[] = [];
  private loading = true;

  private mounted() {
    this.search();
  }

  private search() {
    this.loading = true;
    fetchJSON<DomainResponse>(this.endpoint).then((domainsResponse) => {
      this.loading = false;
      this.domains = domainsResponse.items;
    });
  }
}
</script>
