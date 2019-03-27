<template>
    <div class="col-md-4 col-full">
        <div style="background: #51936e; height: 40px; padding: 4px" class="row justify-content-center">
            <h4 style="color:white">Ruangan</h4>
        </div>
        <div class="row" style="padding-left: 10px; padding-right: 10px; padding-top:7px; padding-bottom:7px;">
            <div class="col-md-12" style="text-align:justify; background:#d2ffe4; border-radius:10px; padding:8px;">
                <div class="container" v-for="(room, index) in ruangans" :key="index">
                    <h6>Ruang: {{room.nama_ruangan}}</h6>
                    <h6>Tipe: {{room.tipe}}</h6>
                    <h6>Jumlah Tersedia: {{room.jumlah}}</h6>
                </div>
            </div>
        </div>
        <div class="row justify-content-md-center">
            <div class="col-md-4">
                <button class="btn btn-sm btn-success">Tambah Ruang</button>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name: 'list-ruangan',
    data() {
        return {
            ruangans: []
        }
    },
    created() {
        const self = this;
        axios.get('http://localhost:8080/room')
            .then(res => {
                console.log(res.data.data);
                self.ruangans = res.data.data;
            })
    },
}
</script>