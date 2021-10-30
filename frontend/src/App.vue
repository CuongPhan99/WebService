<template>
  <div id="app">
    <div class="sidebar" role="banner">
      <div class="header-menu">
        <div class="logo">
          <img src="../src/assets/Mask Group 36.png" alt="" />
        </div>
        <div class="title">
          <img src="../src/assets/best-logos-2017-1.png" alt="" />
          <label>株式会社C-mind</label>
        </div>
        <div class="add-menu">
          <button class="update">
            <img src="../src/assets/Union 6.png" alt="" />
            <p>契約書を作成する</p>
          </button>
        </div>
        <div class="nav-wrap">
          <ul id="nav-sidebar">
            <li v-for="(item, index) in items" v-bind:key="index">
              <a
                @click="
                  index === 3 ? subNav(activeNav) : changeTab();
                  activeClass(index);
                "
                :class="{ active: tabActive == index }"
                >{{ item.name }}
                <img
                  v-if="index === 3"
                  v-bind:src="
                    activeNav
                      ? require('./assets/icon-on.png')
                      : require('./assets/icon-off.png')
                  "
                  alt=""
                />
              </a>
              <ul v-show="activeNav" class="subnav"> 
                <li
                  v-for="(child, index) in item.childrens"
                  v-bind:key="index"
                  @click="activeChildClass(index)"
                  :class="{ activeChild: tabActiveChild == index }"
                >
                  <router-link to="/accounts/1">
                  <a>
                    {{ child.name }}
                  </a>
                  </router-link>

                </li>
              </ul>
            </li>            
          </ul>
        </div>
      </div>
      <div class="footer-menu">
        <ol>
          <li><a href="#">利用規約</a></li>
          <li><a href="#">特定商取引法に基づく表記</a></li>
          <li><a href="#">プライパシーポリシー</a></li>
          <li><a href="#">運営会社</a></li>
          <li><a href="#">よくあるご質問</a></li>
          <li><a href="#">お問い合わせ</a></li>
          <li>
            <ul class="social-network">
              <img src="../src/assets/Mask Group 60.png" alt="" />
              <img src="../src/assets/Mask Group 59.png" alt="" />
              <img src="../src/assets/Mask Group 50.png" alt="" />
              <img src="../src/assets/Mask Group 51.png" alt="" />
            </ul>
          </li>
          <li><p>Digital Sign ©2021</p></li>
        </ol>
      </div>
    </div>
    <div class="content">
      <div class="nav-content">
        <div class="header">
          <div class="row-header">
            <div class="search">
              <img src="../src/assets/Union 5.png" alt="" />
              <input type="text" value="契約書の検索" />
            </div>
            <div class="vip">
              <button class="member-vip">
                <img src="../src/assets/Group 5231.png" alt="" />
                <label>プランを変更</label>
              </button>
            </div>
            <div class="account">
              <img src="../src/assets/1_NpS-s4VDYVsNbIplzmKIMg.png" alt="" />
              <label>水原 コーコー</label>
            </div>
          </div>
        </div>
        <div class="main-title">
          <h4>企業情報</h4>
        </div>
      </div>
      <div class="container">
        <div class="main">
          <AccountCompany v-if="tabActiveChild === 0" />
          <Contact v-if="tabActive === 4" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AccountCompany from "./components/AccountCompany.vue";
import Contact from "./components/Contact.vue";

export default {
  name: "App",
  components: {
    AccountCompany,
    Contact,
  },
  data() {
    return {
      activeNav: false,
      tabActive: "",
      tabActiveChild: "",
      items: [
        { name: "ホーム" },
        { name: "契約書の管理" },
        { name: "雛形の管理" },
        {
          name: ">アカウントの管理",
          childrens: [
            { name: "企業の詳細" },
            { name: "メンバー一覧" },
            { name: "グループ" },
            { name: "請求書" },
          ],
        },
        { name: "顧客一覧" },
      ],
    };
  },
  methods: {
    subNav() {
      this.activeNav = !this.activeNav;
    },
    changeTab() {
      this.activeNav = false;
    },
    activeClass(index) {
      this.tabActive = index;
    },
    activeChildClass(index) {
      this.tabActiveChild = index;
    },
  },
};
</script>

<style>
*,
:before,
:after {
  font-family: "Noto Sans JP";
  box-sizing: border-box;
  outline: none;
  margin: 0;
  padding: 0;
}
/* Side Bar */

.sidebar {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  background: #201a39;
}

.logo {
  display: flex;
  height: 80px;
}
.logo img {
  margin: 25px 70px 25px 30px;
}

.title {
  display: flex;
  height: 60px;
  background-color: #0c0624;
}
.title > img {
  margin: 10px 10px 10px 30px;
  border: 1px solid #ffffff;
  border-radius: 50%;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
}
.title label {
  margin: auto 0;
  color: #707070;
  font-weight: 700;
  font-size: medium;
}

.update {
  border-radius: 5px;
  font-size: 14px;
  width: 210px;
  height: 50px;
  color: white;
  background-color: #d64d10;
  margin: 30px 30px 0px 30px;
  cursor: pointer;
  border: none;
  font-weight: 700;
  font-size: medium;
  position: relative;
}
.update p {
  position: absolute;
  margin: 12px auto;
  left: 50px;
  top: 0;
}
.update img {
  margin: 15px;
  float: left;
}
/* Title */
.main-title {
  padding: 20px 0 20px 30px;
  background-color: #f5f5f5;
}
.main-title h4 {
  margin: 0;
  font-size: 14px;
  color: #271b5a;
}
/* Header */
.nav-content {
  position: fixed;
  z-index: 98;
  top: 0;
  left: 270px;
  right: 0;
  background-color: white;
}
.header {
  padding: 20px 30px 20px 23px;
}
.row-header {
  display: flex;
  height: 40px;
}
.search {
  height: 40px;
  position: relative;
  flex: 1;
}
.search img {
  margin: 12px auto;
}
.search input {
  width: 88px;
  border: none;
  position: absolute;
  top: 11px;
  left: 20px;
  color: #d64d10;
  font-weight: 500;
}
.vip {
  flex: 0 0 135px;
}
.vip button {
  background-color: #d64d10;
  border: none;
  border-radius: 3px;
  height: 40px;
  width: 135px;
  display: flex;
}
.member-vip {
  color: white;
  position: relative;
}
.member-vip > img,
.member-vip > label {
  margin: auto;
  font-size: 14px;
  font-weight: 500;
  font-family: "Noto Sans JP";
}
.account {
  flex: 0 0 141px;
  display: flex;
  margin-left: 30px;
}
.account label {
  margin: 10px 0 10px 10px;
  font-size: 14px;
}
/* Main */
.main {
  margin: 170px 30px 30px 30px;
}

/* Tree Menu */
#nav-sidebar {
  margin: 0;
  padding: 0 30px;
}
#nav-sidebar li {
  list-style-type: none;
}
#nav-sidebar > li:nth-child(4) {
  justify-content: center;
}
#nav-sidebar li a {
  display: flex;
  color: #ffffff;
  text-decoration: none;
  line-height: 40px;
  font-size: 14px;
  font-weight: 700;
}
#nav-sidebar li a img {
  margin: auto;
  margin-right: 0;
}
#nav-sidebar .subnav {
  padding-left: 20px;
}
#nav-sidebar li a.active {
  color: #d64d10;
}
.activeChild {
  position: relative;
  margin: 0 -30px 0 -50px;
  padding-left: 50px;
  background-color: #707070;
  border-left: 4px solid #d64d10;
}

/* Footer Menu */
.footer-menu ol {
  list-style: none;
  margin: 30px auto;
  padding-left: 30px;
}
.footer-menu ol li a {
  font-size: 12px;
  font-weight: 500;
  text-decoration: none;
  color: white;
}
.footer-menu ol li p {
  margin: 0;
  font-size: 10px;
  font-weight: 500;
  color: rgb(179, 171, 171);
}
.social-network {
  display: flex;
  padding: 0;
}
.social-network img {
  text-decoration: none;
  margin: 18px 30px 20px 0px;
}
/* Content */
.content {
  margin-left: 270px;
  flex: 1;
}
</style>
