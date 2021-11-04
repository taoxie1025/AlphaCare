<template>
  <section class="faqs-wrap p-t-b-80" id="faqs-wrap">
    <b-container>
      <div class="row"></div>

      <b-row>
        <b-col md="6" sm="12" class="mb-4 mb-lg-0 mb-md-0 mb-4">
          <div class="section-header mb-5">
            <h2 class="font-weight-bold">FAQs</h2>
            <p></p>
          </div>
          <div v-for="(faq, i) in faqs" role="tablist" v-bind:key="i">
            <b-card no-body class="mb-4 o-hidden">
              <b-card-header header-tag="header" class="p-3" role="tab">
                <b-button
                  block
                  href="#"
                  v-b-toggle="`accordion-` + i"
                  variant="bg-transparent"
                  class="font-weight-700 text-14"
                  >{{ faq.question }}</b-button
                >
              </b-card-header>
              <b-collapse
                :id="`accordion-` + i"
                accordion="my-accordion"
                role="tabpanel"
                :visible="faq.visible"
              >
                <b-card-body>
                  <b-card-text>{{ faq.answer }}</b-card-text>
                </b-card-body>
              </b-collapse>
            </b-card>
          </div>
        </b-col>
        <!-- end of col -->
        <b-col md="5" class="offset-md-1 col-sm-12 d-flex align-items-center">
          <img
            src="@/assets/images/landing/svg/faqs2.svg"
            class="img-responsive w-100"
            alt
          />
        </b-col>
      </b-row>
    </b-container>
  </section>
</template>
<script>
export default {
  data() {
    return {
      faqs: [
        {
          question: "What is Nyx and how does it work?",
          answer:
            "Nyx is a platform that lets any healthcare provider bring their healthcare online in minutes! Create, customize and manage your healthcare portal in simple clicks.",
        },
        {
          question: "Do I get my own website for my healthcare?",
          answer:
            "Yes, immediately after register, you get your own personal healthcare website ready to be customized.",
        },
        {
          question: "What is Nyx app library?",
          answer:
            "Nyx app library is similar to Google or iOS app store. Nyx app library provides rich set of healthcare tools, such as appointments, billing, accounting, prescription, EHR, analytics.",
        },
        {
          question: "How do I customize my healthcare portal?",
          answer:
            "You can browse and search Nyx App Library. To install an app, simply click the install button on the app icon, and it will be immediately available in your portal.",
        },
        {
          question: "Do I need to be technical expert to use Nyx?",
          answer: "No, you can build your healthcare portal similar to .",
        },
        {
          question:
            "Do I need a credit card to try it out and see if it suits my business?",
          answer:
            "Absolutely no. You can upgrade or downgrade your subscription plan at any time and you will be charged on a prorated basis. Our free plan will be free forever and there is no trail period, and your monthly quota will be renewed every 30 days.",
          visible: false,
        },
      ],
    };
  },
};
</script>
<style scoped>
</style>
