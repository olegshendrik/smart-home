angular
.module('angular-map')
.factory 'Fullscreen', ['log', (log) ->
    class Fullscreen
      container: null
      scope: null
      available: false
      callback: null
      isFull: false

      constructor: (@container, @scope, @callback)->
        @init()

      init: ()=>
        @available = document.fullscreenEnabled ||
            document.webkitFullscreenEnabled ||
            document.mozFullScreenEnabled ||
            document.msFullscreenEnabled

      toFull: ()=>
        if @container[0].requestFullscreen
          @container[0].requestFullscreen()
        else if @container[0].webkitRequestFullscreen
          @container[0].webkitRequestFullscreen()
        else if @container[0].mozRequestFullScreen
          @container[0].mozRequestFullScreen()
        else if @container[0].msRequestFullscreen
          @container[0].msRequestFullscreen()

      toMin: ()=>
        if document.cancelFullscreen
          document.cancelFullScreen()
        else if document.webkitCancelFullScreen
          document.webkitCancelFullScreen()
        else if document.mozCancelFullScreen
          document.mozCancelFullScreen()

      status: ()->
        window.innerHeight == @container.height()

      resize: ()=>
        if @status()
          @toMin()
        else
          @toFull()

    Fullscreen
]