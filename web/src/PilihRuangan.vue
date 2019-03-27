<template>
    <div class="col-md-8 col-full">
        <div style="background:#5daf82; height: 40px; padding: 5px" class="row justify-content-md-center">
            <div class="col-md-3">
                <h4 style="color:white">Pesan Ruangan</h4>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6" style="padding: 10px; border-bottom: 1px solid #c3c3c3;">
                <button class="btn btn-sm btn-primary" @click="updateStep(2)">Kembali</button>
            </div>
        </div>
        <div class="row justify-content-md-center" style="background: #dfdfdf; padding-top: 10px; padding-bottom: 10px;">
            <div class="col-md-12">
                <div class="input-group mb-3" style="padding: 10px;">
                    <div class="input-group-prepend">
                        <span class="input-group-text" id="input-nama">Pasien</span>
                    </div>
                    <input type="text" v-model="pasien.nama_pasien" class="form-control" placeholder="eg. John Doe" aria-label="Nama" aria-describedby="input-nama" disabled>
                </div>
                <div class="input-group mb-3" style="padding: 10px;">
                    <div class="input-group-prepend">
                        <span class="input-group-text" id="input-tujuan">Tujuan</span>
                    </div>
                    <input type="text" v-model="rs.nama_rs" class="form-control" placeholder="eg. RS Apa" aria-label="Tujuan" aria-describedby="input-tujuan" disabled>
                </div>
            </div>
        </div>
        <template v-for="(room, index) in rooms">
        <div v-if="index % 2 == 0" class="row justify-content-md-center" style="background: #d2ffe4; padding-top: 10px; padding-bottom: 10px;">
            <div class="col-md-10">
                <h5>Ruang: {{room.nama_ruangan}}</h5>
                <h5>Tipe: {{room.tipe}}</h5>
                <h5>Tersedia: {{room.jumlah}}</h5>
            </div>
            <div class="col-md-2 align-self-center">
                <button class="btn btn-sm btn-primary" @click="pesan(room)">Pesan</button>
            </div>
        </div>
        <div v-else class="row justify-content-md-center" style="background: #dfdfdf; padding-top: 10px; padding-bottom: 10px;">
            <div class="col-md-10">
                <h5>Ruang: {{room.nama_ruangan}}</h5>
                <h5>Tipe: {{room.tipe}}</h5>
                <h5>Tersedia: {{room.jumlah}}</h5>
            </div>
            <div class="col-md-2 align-self-center">
                <button class="btn btn-sm btn-primary" @click="pesan(room)">Pesan</button>
            </div>
        </div>
        </template>
    </div>
</template>
<script>
export default {
    name: 'pilih-pasien',
    props: ["pasien", "rs"],
    data() {
        return {
            rooms: []
        }
    },
    methods: {
        pesan(item) {
            const self = this;
            axios.post('http://localhost:8080/network/transfer', {
                ip: self.rs.ip,
                id: item.id,
                pasien: {
                    nama_pasien: self.pasien.nama_pasien,
                    diagnosa: self.pasien.diagnosa,
                }
            })
            .then(res => {
                self.updateStep(0);
            })
            .catch(err => console.log(err))
        },
        updateStep(index) {
            this.$emit('chstep', index)
        },
    },
    created() {
        const self = this;
        axios.post('http://localhost:8080/network/rooms', {
            ip: self.rs.ip,
        })
        .then(res => {
            self.rooms = res.data.data;
        })
    },
}
</script>