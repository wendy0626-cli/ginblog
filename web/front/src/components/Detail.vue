<template>
  <v-card min-height="100vh">
    <!-- 文章标题 -->
    <div class="pa-3 text-center text-h5 font-weight-bold blue--text text--lighten-2">
      {{artInfo.title}}
    </div>
    <!-- <v-divider class="ma-3 pa-2"></v-divider> -->
    <div class="ml-3 mr-3 mt-4 pl-2 pr-3">
      <v-icon small>{{'mdi-account'}}</v-icon>
      <span class="text-caption">作者：</span>
      <span class="text-caption">温迪</span>

      <v-icon small class="ml-5">{{'mdi-calendar-month'}}</v-icon>
      <span class="text-caption">创建于</span>
      <span class="text-caption">{{artInfo.CreatedAt | dateformat('YYYY-MM-DD HH:SS')}}</span>

      <v-icon small class="ml-5">{{'mdi-calendar-month'}}</v-icon>
      <span class="text-caption">更新于</span>
      <span class="text-caption">{{artInfo.UpdatedAt | dateformat('YYYY-MM-DD HH:SS')}}</span>
    </div>
    <div class="d-flex mr-10 justify-center">
      <v-icon class="mr-1" color="green" small>{{ 'mdi-eye' }}</v-icon>
      <span>{{ artInfo.read_count }}</span>
    </div>
    <v-divider class="ma-3 pa-2"></v-divider>

    <!-- 概述 -->
    <v-alert border="left" elevation="2" color="blue" dark outlined class="ma-4 pa-3">
      {{artInfo.desc}}
    </v-alert>
    <!-- 文章内容 -->
    <div class="content ma-4 pa-3 text-justify" v-html="artInfo.content"></div>

    <!-- 评论内容-->
    <v-divider class="ma-5"></v-divider>
    <v-sheet class="ma-3 pa-3">
      <div>
        <v-list
          outlined
          class="ma-3 pa-3"
          v-for="item in commentList"
          :key="item.ID"
          v-show="item.status === 1"
        >
          <template>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>
                  {{ item.username }}
                  {{
                  item.CreatedAt | dateformat('YYYY-MM-DD')
                  }}
                </v-list-item-title>
                <v-list-item-subtitle class="mr-3">
                  {{
                  item.content
                  }}
                </v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </template>
        </v-list>
      </div>
      <!-- <div class="text-center" v-if="commentList">
        <v-pagination
          class="my-2"
          total-visible="7"
          v-model="queryParam.pagenum"
          :length="Math.ceil(total / queryParam.pagesize)"
          @input="getCommentList()"
        ></v-pagination>
      </div> -->
    </v-sheet>

  </v-card>

   

</template>
<script>
import Prism from 'prismjs';

export default {
  props: ['id'],
  data () {
    return {
      artInfo: {},
      commentList: [],
      queryParam: {
        pagesize: 1,
        pagenum: 1
      }
    }
  },
  created () {
    this.getArtInfo();
    this.getCommentList();
  },
  methods: {
    //查询文章
    async getArtInfo () {
      const { data: res } = await this.$axios.get(`article/${this.id}`);
      this.artInfo = res.data;
      window.sessionStorage.setItem('title', this.artInfo.title)
    },
    //获取评论
    async getCommentList() {
      const { data: res } = await this.$axios.get(`commentfront/${this.id}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.commentList = res.data
      this.total = res.total
    }
  },
}
</script>
<style scoped>
.content >>> img,
span,
p {
  max-width: 96%;
}
</style>