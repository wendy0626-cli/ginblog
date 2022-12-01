<template>
	<div>
		<a-card>
			<h3>个人设置</h3>
			<a-form-model>
				<!-- 作者名称 -->
				<a-form-model-item label="作者名称" prop="name">
					<a-input v-model="profile.name"></a-input>
				</a-form-model-item>
				<!-- 座右铭 -->
				<a-form-model-item label="座右铭" prop="motto">
					<a-input v-model="profile.motto"></a-input>
				</a-form-model-item>
				<!-- 个人介绍 -->
				<a-form-model-item label="个人介绍" prop="desc">
					<a-textarea v-model="profile.desc"></a-textarea>
				</a-form-model-item>
				<!-- QQ -->
				<a-form-model-item label="QQ" prop="qq_chat">
					<a-input v-model="profile.qq_chat"></a-input>
				</a-form-model-item>
				<!-- 微信 -->
				<a-form-model-item label="微信" prop="wechat">
					<a-input v-model="profile.wechat"></a-input>
				</a-form-model-item>
				<!-- bilibili -->
				<a-form-model-item label="哔哩哔哩" prop="bili">
					<a-input v-model="profile.bili"></a-input>
				</a-form-model-item>
				<!-- 码云 -->
				<a-form-model-item label="码云" prop="gitee">
					<a-input v-model="profile.gitee"></a-input>
				</a-form-model-item>
				<!-- 邮箱 -->
				<a-form-model-item label="邮箱" prop="email">
					<a-input v-model="profile.email"></a-input>
				</a-form-model-item>
				<!-- 背景图 -->
				<a-form-model-item label="背景图" prop="img">
					<a-input v-model="profile.img"></a-input>
				</a-form-model-item>
				<!-- 头像 -->
				<a-form-model-item label="头像" prop="avatar">
					<a-input v-model="profile.avatar"></a-input>
				</a-form-model-item>
				<!-- 提交 -->
				<a-form-model-item>
					<a-button type="danger" style="margin-right:15px" @click="updateProfile">更新</a-button>
				</a-form-model-item>
			</a-form-model>
		</a-card>
	</div>
</template>
<script>
	export default {
		data() {
			return {
				profile: {
					id: 1,
					name: '',
					motto: '',
					desc: '',
					qq_chat: '',
					wechat: '',
					bili: '',
					gitee: '',
					email: '',
					img: '',
					avatar: ''
				}
			}
		},
		created() {
			this.getProfileInfo();
		},
		methods: {
			async getProfileInfo() {
				const { data: res } = await this.$axios.get(`profile/${this.profile.id}`);
				if (res.status !== 200) return this.$message.error(res.message)
				this.profile = res.data
			},
			// 更新
			async updateProfile() {
				const { data: res } = await this.$axios.put(`profile/${this.profile.id}`, this.profile)
				if (res.status !== 200) return this.$message.error(res.message)
				this.$message.success(`个人信息更新成功`)
				this.$router.push('/index')
			},
		},
	}
</script>
<style>

</style>