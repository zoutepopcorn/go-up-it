<template>
  <div>
    <v-snackbar v-model="oops" timeout="2000" color="error" top outlined>
      Oops cant upload the files
    </v-snackbar>

    <v-file-input v-model="files" ref="files" placeholder="Upload your documents" label="File input"
                  @change="updateFiles" counter outlined multiple show-size prepend-icon="mdi-paperclip"
                  :disabled="isUploading">
      <template v-slot:selection="{ text }">
        <v-chip small label color="primary">
          {{ text }}
        </v-chip>
      </template>
    </v-file-input>
    <v-btn @click="submitFilesByOne()" color="blue" dark block elevation="2" :disabled="isUploading">
      <v-icon>mdi-send</v-icon>
      upload
    </v-btn>
    <br>
    <v-card flat>
      <div v-for="(file, i) of fileInfo">
        <v-chip class="mb-1 mt-2" color="blue darken-2" text-color="white" label small>
          <v-icon small left>
            mdi-paperclip
          </v-icon>
          {{ file.name }}
        </v-chip>
        <v-progress-linear v-model="file.uploadPercentage" color="blue" height="19" dark>
          <strong>{{ file.uploadPercentage }} {{ file.uploadPercentage ? "%" : "" }}</strong>
        </v-progress-linear>
      </div>
    </v-card>
    <!--    <v-progress-linear v-model="uploadPercentage" color="blue" height="19" dark>-->
    <!--      <strong>{{ uploadPercentage }} {{ uploadPercentage[i] ? "%" : "" }}</strong>-->
    <!--    </v-progress-linear>-->
    <!--    <v-progress-linear v-if="uploadPercentage[i] > 0 && uploadPercentage[i] < 100" color="deep-purple accent-4" indeterminate-->
    <!--                       rounded height="1"></v-progress-linear>-->
  </div>
</template>

<script>
import axios from 'axios'

const ORIGIN = window.origin;

export default {
  data() {
    return {
      files: [],
      fileInfo: [],
      uploadPercentage: [],
      isUploading: false,
      oops: false
    }
  },
  created() {
    // setInterval(() => {
    //     this.snackbar = !this.snackbar;
    // }, 2000)
  },
  computed: {},
  methods: {
    updateFiles() {
      this.fileInfo = [];
      for (const file of this.files) {
        const ITEM = {name: file.name, size: file.size, uploadPercentage: 0};
        this.fileInfo.push(ITEM);
      }
    },
    async sendData() {
      // console.log(this.$refs);
      // const file = event.target.files[0]
      // axios.post('upload_file', file, {
      //   headers: {
      //     'Content-Type': file.type
      //   }
      // })
      // console.log();
    },
    async submitFilesByOne() {
      let i = 0;
      this.isUploading = true;
      for (const file of this.files) {
        const formData = new FormData();
        formData.append("files", file, file.name);
        try {
          const onUploadProgress = (event) => {
            const percentCompleted = Math.round((event.loaded * 100) / event.total);
            this.fileInfo[i].uploadPercentage = percentCompleted
          };

          const response = await axios.post(`${ORIGIN}/multi`, formData, {onUploadProgress})
          console.log({response});
        } catch (error) {
          console.log({error});
          this.fileInfo[i].uploadPercentage = 0
          this.oops = true;
        }
        i++;
      }
      this.isUploading = false;
    },
    async submitFiles() {
      if (this.files) {
        const formData = new FormData();
        for (let file of this.files) {
          formData.append("files", file, file.name);
        }
        try {
          const onUploadProgress = event => {
            console.log("progress", event)
            const percentCompleted = Math.round((event.loaded * 100) / event.total);
            console.log('onUploadProgress', percentCompleted);
            this.uploadPercentage = percentCompleted;
          };

          const response = await axios.post(`${ORIGIN}/multi`, formData, {onUploadProgress})
          console.log({response});
        } catch (error) {
          console.log({error});
        }
      } else {
        console.log("there are no files.");
      }
    }
  }
}
</script>