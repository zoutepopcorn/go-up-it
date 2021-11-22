<template>
  <v-app>
    <v-main>

      <v-dialog v-model="isQr" width="500">
        <v-card>
          <v-card-title class="text-h5 grey lighten-2">
            Open on device
          </v-card-title>

          <v-card-text class="text-center">
            <img :src="qrSource" alt=""><br>
            {{url}}
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="isQr = false">
              Close
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-app-bar color="white" elevation="3">
        <v-icon color="blue">mdi-paperclip</v-icon>
        <h1 class="ml-3 overline" color="blue--text"> File uploader</h1>
      </v-app-bar>
      <v-spacer></v-spacer>
      <v-card max-width="600" class="pa-8 ma-8 mx-auto">
        <UploadForm/>
      </v-card>
      <v-spacer></v-spacer>
    </v-main>
  </v-app>
</template>

<script>
import UploadForm from './components/UploadForm';
// import bwipjs from 'bwip-js';
import QRCode from 'qrcode'

export default {
  name: 'App',

  components: {
    UploadForm,
  },
  async created() {
    const params = new URLSearchParams(window.location.search)
    this.isQr = !!params.get('qr');
    console.log(this.isQr);
    if(this.isQr) {
      try {
        const MY_IP_REQ = await fetch('/api/ip');
        const MY_IP = await MY_IP_REQ.text();
        const URL = `http://${MY_IP}:3000/`;
        this.url = URL;
        this.qrSource = await QRCode.toDataURL(URL, {scale: 8});
      } catch (err) {
        console.error(err)
      }
    }
  },
  data: () => ({
    isQr: false,
    qrSource: "test",
    url: ""
  }),
};
</script>
