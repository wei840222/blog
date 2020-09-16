<template>
  <div id="app">
    <ApolloQuery :query="require('./graphql/posts.gql')">
      <template slot-scope="{ result: { loading, error, data } }">
        <div v-if="loading" class="loading apollo">Loading...</div>
        <div v-else-if="error" class="error apollo">An error occured</div>
        <div v-else-if="data" class="result apollo">
          <div v-for="post in data.posts" :key="post.id">
            <el-card class="box-card" shadow="hover" style="min-height: 85px;">
              <el-row>
                <el-col :span="4">
                  <el-row>
                    <el-avatar size="small" :src="post.user.avatarUrl" />
                  </el-row>
                  <el-row>{{post.user.name }}</el-row>
                </el-col>
                <el-col :span="20" style="text-align: left;">{{post.text }}</el-col>
              </el-row>
              <el-row>
                <el-divider />
              </el-row>
              <el-row>{{post.comments }}</el-row>
            </el-card>
          </div>
        </div>
        <div v-else class="no-result apollo">No result :(</div>
      </template>
    </ApolloQuery>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.box-card {
  margin-bottom: 20px;
}
</style>
