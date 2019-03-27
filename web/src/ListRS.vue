<template>
    <div class="col-md-8 col-full">
        <div style="background:#5daf82; height: 40px; padding: 5px" class="row justify-content-md-center">
            <div class="col-md-4">
                <h4 style="color:white">Daftar Rumah Sakit</h4>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6" style="padding: 10px; border-bottom: 1px solid #c3c3c3;">
                <button class="btn btn-sm btn-primary" @click="updateStep(0)">Kembali</button>
            </div>
        </div>
        <div class="row justify-content-md-center" style="background: #dfdfdf; padding-top: 10px; padding-bottom: 10px;">
            <div class="col-md-12">
                <div class="input-group mb-3" style="padding: 10px;">
                    <div class="input-group-prepend">
                        <span class="input-group-text" id="input-nama">Pasien</span>
                    </div>
                    <input type="text" class="form-control" v-model="pasien.nama_pasien" placeholder="eg. John Doe" aria-label="Nama" aria-describedby="input-nama" disabled>
                </div>
            </div>
        </div>
        <template v-for="(item, index) in list">
            <div v-if="index % 2 == 0" class="row justify-content-md-center" style="background: #dfdfdf; padding-top: 10px; padding-bottom: 10px;">
                <div class="col-md-10">
                    <h5>{{item.nama_rs}}</h5>
                </div>
                <div class="col-md-2 align-self-center">
                    <button class="btn btn-sm btn-primary" @click="send(item)">Lihat Ruangan</button>
                </div>
            </div>
            <div v-else class="row justify-content-md-center" style="background: #d2ffe4; padding-top: 10px; padding-bottom: 10px;">
                <div class="col-md-10">
                    <h5>{{item.nama_rs}}</h5>
                </div>
                <div class="col-md-2 align-self-center">
                    <button class="btn btn-sm btn-primary" @click="send(item)">Lihat Ruangan</button>
                </div>
            </div>
        </template>
    </div>
</template>
<script>
export default {
    name: 'list-rs',
    props: ["pasien"],
    data() {
        return {
            list: []
        }
    },
    methods: {
        updateStep(index) {
            this.$emit('chstep', index)
        },
        send(item) {
            this.$emit('rs', item)
        }
    },
    created() {
        const self = this;
        axios.get('http://localhost:8080/network')
        .then(res => {
            self.list = res.data.data;
        })
    },
}
</script>