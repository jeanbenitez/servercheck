<template>
  <div>
    <b-form id="form-domain" inline @submit="onSubmit">
      <b-form-group
        class="fieldset-domain"
        description="Type a valid domain."
        label="Domain"
        label-for="input-domain"
        :invalid-feedback="invalidFeedback"
        :valid-feedback="validFeedback"
        :state="isValidDomain"
      >
        <b-input-group>
          <b-form-input id="input-domain" v-model="domain" :state="isValidDomain" trim></b-form-input>
          <b-input-group-append>
            <b-button variant="outline-secondary" type="submit">Search</b-button>
          </b-input-group-append>
        </b-input-group>
      </b-form-group>
    </b-form>

    <DomainTable :domains="domainData" :loading="loading" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Model } from 'vue-property-decorator';
import { fetchJSON } from '../utils/http';
import { Domain } from '../interfaces/domain';
import DomainTable from './DomainTable.vue';

@Component({
  components: {
    DomainTable,
  },
})
export default class SearchDomain extends Vue {
  public domain = 'truora.com';
  public domainData: Domain[] = [{} as Domain];
  public loading = true;

  private mounted() {
    this.search();
  }

  get invalidFeedback() {
        if (this.isValidDomain) {
          return '';
        } else if (this.domain && this.domain.length > 0) {
          return 'Enter a valid domain';
        } else {
          return 'Please enter something';
        }
  }

  get validFeedback() {
        if (this.isValidDomain) {
          return 'Ok';
        }
  }

  get isValidDomain() {
    return /^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/g.test(this.domain);
  }

  get endpoint() {
    return 'http://localhost:8005/domains/' + this.domain;
  }

  private onSubmit(evt) {
    evt.preventDefault();
    this.search();
  }

  private search() {
    if (!this.isValidDomain) {
      return;
    }

    this.loading = true;
    fetchJSON<Domain>(this.endpoint).then((domain) => {
        this.loading = false;
        this.domainData = [domain];
    });
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#form-domain {
  max-width: 500px;
  width: 100%;
  margin: 30px auto;
}
#form-domain .fieldset-domain {
  flex-flow: column wrap;
  flex-direction: column;
  text-align: left;
  align-items: start;
}
</style>
