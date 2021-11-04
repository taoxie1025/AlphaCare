<template>
  <!-- start work carousel -->
  <section id="works-wrap" class="work-carousel-wrap">
    <b-container>
      <b-row>
        <div class="section-header col-md-12 mb-5">
          <h2 class="font-weight-bold">
            Your own healthcare platform in few easy steps
          </h2>
        </div>
        <b-col md="12">
          <div class="slick-work-carousel-wrap">
            <!-- carousel item start -->
            <carousel
              :autoplay="false"
              :dots="false"
              :nav="false"
              :responsive="{
                0: { items: 1, nav: true, navText: false },
                1000: { items: 3, nav: true, navText: false },
              }"
              :slideBy="3"
            >
              <div v-for="(step, key) in steps" :key="key">
                <div class="card text-left o-hidden mb-4 m-3">
                  <img
                    class="card-img-top"
                    :src="step.path"
                    alt="Card image cap"
                  />
                  <div class="card-body">
                    <h5 class="card-title font-weight-bold">
                      {{ step.title }}
                    </h5>
                    <p class="card-text">
                      {{ step.content }}
                    </p>
                  </div>
                </div>
              </div>
            </carousel>
          </div>
        </b-col>
      </b-row>
    </b-container>
  </section>
  <!-- end work carousel -->
</template>
<script>
import carousel from "vue-owl-carousel";
export default {
  components: { carousel },
  data() {
    return {
      steps: [
        {
          path: require("@/assets/images/landing/products/products12.svg"),
          key: "step-1",
          title: "Register",
          content:
            "Sign up for a free account. Get your own healthcare portal immediately.",
        },
        {
          path: require("@/assets/images/landing/products/products11.svg"),
          key: "step-2",
          title: "Build and Customize",
          content: "One click install any app/tool from our library.",
        },
        {
          path: require("@/assets/images/landing/products/products2.svg"),
          key: "step-3",
          title: "Patient portal",
          content:
            "Add your patients. Patients grant access to powerful tools like EHR, appointments, billing through patient portal.",
        },
      ],
    };
  },
};
</script>
