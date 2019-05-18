<template>
  <div class="domain-table">
    <Loading v-if="loading" />
    <b-table striped :items="formattedDomains">
      <template slot="servers" slot-scope="servers">
        <b-table hover :items="formatItems(servers.value)"></b-table>
      </template>
      <template slot="logo" slot-scope="logo">
        <img :src="logo.value" class="logo" />
      </template>
    </b-table>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Domain } from '../interfaces/domain';
import Loading from './Loading.vue';

@Component({
  components: {
    Loading,
  },
})
export default class DomainTable extends Vue {
  @Prop() public domains: Domain | Domain[];
  @Prop() public loading: boolean;

  get formattedDomains(): Domain[] {
    return this.formatItems(this.domains);
  }

  private formatItems(items) {
    return Array.isArray(items)
      ? items.filter(Boolean)
      : items
        ? [items]
        : [];
  }
}
</script>

<style scoped>
.domain-table {
  width: 95%;
  margin: 20px auto;
}
.domain-table .logo {
  max-width: 100%;
  margin: 0 auto;
  display: block;
}
</style>
