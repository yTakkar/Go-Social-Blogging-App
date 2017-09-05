import $ from 'jquery'
import Notify from 'handy-notification'
import * as fn from './utils/functions'

$(() => {

  // SIGNUP FUNCTION
  $('form.form_register').submit(e => {
    e.preventDefault()

    let
      username = $('.r_username').val(),
      email = $('.r_email').val(),
      password = $('.r_password').val(),
      password_again = $('.r_password_again').val()

    if (!username || !email || !password || !password_again) {
      Notify({ value: 'Values are missing!' })
    } else if (password != password_again) {
      Notify({ value: 'Passwords don\'t match!' })
    } else {

      let signupOpt = {
        data: {
          username,
          email,
          password,
          password_again
        },
        btn: $('.r_submit'),
        url: '/user/signup',
        redirect: '/',
        defBtnValue: 'Sign up for free',
      }
      fn.commonLogin(signupOpt)
    }

  })

  // LOGIN FUNCTION
  $('form.form_login').submit(e => {
    e.preventDefault()

    let
      username = $('.l_username').val(),
      password = $('.l_password').val()

    if (!username || !password) {
      Notify({ value: 'Values are missing!' })
    } else {

      let loginOpt = {
        data: {
          username: $('.l_username').val(),
          password: $('.l_password').val()
        },
        btn: $('.l_submit'),
        url: '/user/login',
        redirect: '/',
        defBtnValue: 'Login to continue',
      }
      fn.commonLogin(loginOpt)

    }

  })

  // CREATE BLOG
  $('.create_blog').on('click', e => {
    e.preventDefault()
    let
      title = $('.e_username').val(),
      content = $('.e_bio').val()

    if(!title || !content){
      Notify({ value: 'Some values are missing!!' })
    } else {
      $.ajax({
        url: '/api/create-new-blog',
        data: {
          title,
          content
        },
        method: 'POST',
        dataType: 'JSON',
        success: data => {
          let { mssg } = data
          Notify({ value: mssg, done: () => location.href = '/' })
        }
      })
    }

  })

  // SCROLL TO TOP
  $('.page_end').on('click', () => $('html, body').animate({ scrollTop: 0 }, 450) )

  // VIEW BLOG FUNCTIONS
  // FN FOR GOING BACK FROM VIEWING BLOG
  $('.blog_back').on('click', e => fn.Back(e) )

  // FOR DELETING THE BLOG
  $('.blog_dlt').on('click', e => {
    e.preventDefault()
    let blog = e.currentTarget.dataset['blog']
    $.ajax({
      url: '/api/delete-blog',
      data: { blog },
      method: 'POST',
      dataType: 'JSON',
      success: data => {
        Notify({ value: data.mssg, done: () => location.href = '/' })
      }
    })
  })

  // FOR EDITING THE BLOG
  $('.d_editing').on('click', e => {
    e.preventDefault()
    let
      blogID = $('.edit').data('blog'),
      title = $('.b_title').val(),
      content = $('.b_content').val()

    if (!title || !content){
      Notify({ value: 'Some values are missing!!' })
    } else {
      $.ajax({
        url: '/api/edit-blog',
        data: {
          blogID,
          title,
          content
        },
        method: 'POST',
        dataType: 'JSON',
        success: data => {
          Notify({ value: data.mssg, done: () => location.reload() })
        }
      })
    }
  })

})
