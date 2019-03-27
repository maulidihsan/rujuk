<template>
    <div class="col-md-8 col-full">
        <div style="background:#5daf82; height: 40px; padding: 5px" class="row justify-content-md-center">
            <div class="col-md-3">
                <h4 style="color:white">Input Pasien</h4>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6" style="padding: 10px; border-bottom: 1px solid #c3c3c3;">
                    <button class="btn btn-sm btn-primary" @click="updateStep(0)">Kembali</button>
            </div>
            <div class="col-md-6 text-right" style="padding: 10px; border-bottom: 1px solid #c3c3c3;">
                    <button class="btn btn-sm btn-success" @click="save">Simpan</button>
            </div>
        </div>
        <div class="row justify-content-md-center" style="background: #dfdfdf; padding-top: 10px; padding-bottom: 10px;">
            <div class="col-md-12">
                <form>
                    <div class="input-group mb-3" style="padding: 10px;">
                        <div class="input-group-prepend">
                            <span class="input-group-text" id="input-nama">Pasien</span>
                        </div>
                        <input v-model="nama" type="text" class="form-control" placeholder="eg. John Doe" aria-label="Nama" aria-describedby="input-nama">
                    </div>
                    <div class="input-group mb-3" style="padding: 10px;">
                        <div class="input-group-prepend">
                            <span class="input-group-text">Diagnosa   </span>
                        </div>
                        <textarea v-model="diagnosa" class="form-control" aria-label="Diagnosa"></textarea>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name: 'input-pasien',
    data() {
        return {
            nama: '',
            diagnosa: ''
        }
    },
    methods: {
        updateStep(index) {
            this.$emit('chstep', index)
        },
        save() {
            const self = this;
            axios.post('http://localhost:8080/pasien', {
                nama_pasien: this.nama,
                diagnosa: this.diagnosa
            })
            .then(res => {
                console.log(res.data);
                self.updateStep(0);
            })
        }
    },
}
</script>