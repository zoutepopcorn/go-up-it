<template>
  <v-app>
    <v-main>
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

export default {
  name: 'App',

  components: {
    UploadForm,
  },

  data: () => ({
    //
  }),
};
</script>
