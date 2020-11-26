<template>
  <v-app>
    <v-app-bar
      app
      color="teal"
      dark
    >
      <v-app-bar-nav-icon
        @click="drawer = !drawer"
        color="black"
      ></v-app-bar-nav-icon>
      <v-toolbar-title>Snake Game</v-toolbar-title>
      <v-spacer></v-spacer>
        <span class="mr-2">Made by Oscar Andres Montes</span>
        <v-icon>mdi-alpha-c-circle-outline</v-icon>
        <v-icon>mdi-alien</v-icon>
    </v-app-bar>
    <v-navigation-drawer
    v-model="drawer"
    app
    absolute
    bottom
    color="accent"
    temporary
    >
      <v-layout mt-5 column align-center>
        <v-flex>
        <v-avatar
          size="103"
        >
          <v-img :src="require('./assets/Savatar.jpeg')" alt="Snake Avatar"></v-img>
        </v-avatar>
        </v-flex>
        <v-divider></v-divider>
        <v-flex>
          <h2 class="white--text headline mt-5">Snake</h2>
        </v-flex>
        <v-btn class="mt-3"
          elevation="24"
          @click="form = true"
          > Play Game </v-btn>
        <v-divider></v-divider>
        <v-btn class="mt-8"
            elevation="24"
            @click="HighScores"
        > HighScores </v-btn>
      </v-layout>
    </v-navigation-drawer>
    <v-main>
      <v-dialog v-model="scores" width="500">
        <v-card
            elevation="13"
        >
          <v-card-title>Highscores</v-card-title>
          <v-card-text>
            <v-simple-table>
              <template v-slot:default>
                <thead>
                <tr>
                  <th class="text-left">
                    Name
                  </th>
                  <th class="text-left">
                    Score
                  </th>
                </tr>
                </thead>
                <tbody>
                <tr
                    v-for="item in players"
                    :key="item.name"
                >
                  <td>{{ item.name }}</td>
                  <td>{{ item.score }}</td>
                </tr>
                </tbody>
              </template>
            </v-simple-table>
          </v-card-text>
        </v-card>
      </v-dialog>
      <v-dialog v-model="form" width="500">
        <v-card
          elevation="13"
        >
        <v-form @submit.prevent="registerUser">
          <v-text-field class="pl-2 pt-4" label="Player's Name" v-model="player"></v-text-field>
          <v-btn class="mb-2 ml-2"  type="submit" color="success">Enter</v-btn>
          <v-btn v-if="game" class="mb-2 ml-16" color="warning" @click="startGame"> Start </v-btn>
        </v-form>
        </v-card>
      </v-dialog>
      <v-snackbar
          v-model="snackbar"
          :vertical= true
      > Please Insert a name
        <template>
          <v-btn class="ml-16"
              color="pink"
              text
              @click="snackbar = false"
          >
            Close
          </v-btn>
        </template>
      </v-snackbar>
      <Init/>
    </v-main>
    <v-footer
        dark
        padless
    >
      <v-card
          class="flex"
          flat
          tile
      >
        <v-card-title class="teal text-center">
          <v-spacer></v-spacer>
          <strong>Get connected with my social networks!</strong>
          <v-btn
              href="https://twitter.com/damstyx"
              class="mx-4"
              dark
              icon
          >
            <v-icon size="24px">
              mdi-twitter
            </v-icon>
          </v-btn>
          <v-btn
              href="https://www.linkedin.com/in/andmontc/"
              class="mx-4"
              dark
              icon
          >
            <v-icon size="24px">
              mdi-linkedin
            </v-icon>
          </v-btn>
          <v-btn
              href="https://www.facebook.com/Andmont"
              class="mx-4"
              dark
              icon
          >
            <v-icon size="24px">
              mdi-facebook
            </v-icon>
          </v-btn>
        </v-card-title>

        <v-card-text class="py-2 white--text text-center">
          {{ new Date().getFullYear() }} — <strong>Made with Vuetify</strong>
        </v-card-text>
      </v-card>
    </v-footer>
  </v-app>
</template>

<script>
import Init from './components/Init';
import Axios from "axios";


export default {
  name: 'App',

  components: {
    Init,
  },

  data: () => ({
    drawer:false,
    form: false,
    players: [],
    player:'',
    snackbar: false,
    scores: false,
    game:false,
  }),
  methods: {
    async registerUser () {
      if (this.player === '') {
        this.snackbar= true
      } else {
        this.game = true
        await Axios.post('http://localhost:3000/player', {
          name: this.player,
        })
      }
    },
    async HighScores () {
      this.scores = true;
      await Axios.get('http://localhost:3000/players')
      .then(ans =>{
        this.players = ans.data;
      })
      .catch(err => {
        console.log(err);
      })
    },
    startGame() {
      window.location.href = "http://localhost:8000/Snakegame.html"
    }
  }
};
</script>