<template>
<div class="container-fluid">
    <div class="row">
        <template v-if="step == 0">
            <list-pasien @pasien="selectPasien" @chstep="updateStep" />
        </template>
        <template v-else-if="step == 1">
            <input-pasien  @chstep="updateStep"/>
        </template>
        <template v-else-if="step == 2">
            <list-rs :pasien="pasien" @chstep="updateStep" @rs="selectRs" />
        </template>
        <template v-else-if="step == 3">
            <pilih-ruangan  :pasien="pasien" :rs="rs" @chstep="updateStep" />
        </template>
        <list-ruangan />
    </div>
</div>
</template>
<script>
import InputPasien from './InputPasien.vue'
import ListPasien from './ListPasien.vue'
import ListRs from './ListRS.vue'
import ListRuangan from './ListRuangan.vue'
import PilihRuangan from './PilihRuangan.vue'
export default {
    name: 'main-menu',
    components: {
        InputPasien,
        ListPasien,
        ListRs,
        ListRuangan,
        PilihRuangan
    },
    data () {
        return {
            step: 0,
            pasien: {},
            rs: {},
        }
    },
    methods: {
        updateStep(inp) {
            this.step = inp;
        },
        selectPasien(item) {
            console.log(item);
            this.pasien = item;
            this.step = 2;
        },
        selectRs(item) {
            this.rs = item;
            this.step = 3;
        }
    },
}
</script>