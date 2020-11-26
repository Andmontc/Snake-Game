<template>
  <v-card
      class="mx-auto"
      max-width="500"
  >
    <v-img
        height="400"
        :src="require('../assets/logo.png')"
        class="my-3"
        contain
    ></v-img>
    <v-card-title class="pt-0">Snake Game - Instructions</v-card-title>
    <v-card-text>
      <v-row
          align="center"
          class="mx-0"
      >
        <v-rating
            :value="4.5"
            color="amber"
            dense
            half-increments
            readonly
            size="14"
        ></v-rating>

        <div class="grey--text ml-4">
          4.5 (410)
        </div>
      </v-row>
      <div>To start the game press the black lines button at the top left of the screen and press Play Game, to see the scores press the highscores button.</div>
    </v-card-text>
    <v-card-text>
      <v-chip-group>
        <v-chip>Snake</v-chip>
        <v-chip>Game</v-chip>
      </v-chip-group>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'Init',

  data: () => ({
    //
  }),
}
</script>
