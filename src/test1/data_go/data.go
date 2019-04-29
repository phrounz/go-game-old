package data_go

func GetBytesFromFilename(filename string) []byte {
  if len(filename)>=2 && filename[0:2]!="./" { filename = "./"+filename }
  switch (filename) {
  case "./data/level1/0-collision.png": return level1__0_collision_png
  case "./data/level1/0-front.png": return level1__0_front_png
  case "./data/level1/0.png": return level1__0_png
  case "./data/level1/1-collision.png": return level1__1_collision_png
  case "./data/level1/1.png": return level1__1_png
  case "./data/level1/2-collision.png": return level1__2_collision_png
  case "./data/level1/2.png": return level1__2_png
  case "./data/level1/3-collision.png": return level1__3_collision_png
  case "./data/level1/3-front.png": return level1__3_front_png
  case "./data/level1/3.png": return level1__3_png
  case "./data/level1/small/m-105.png": return level1__small__m_105_png
  case "./data/level1/small/m-120.png": return level1__small__m_120_png
  case "./data/level1/small/m-135.png": return level1__small__m_135_png
  case "./data/level1/small/m-15.png": return level1__small__m_15_png
  case "./data/level1/small/m-150.png": return level1__small__m_150_png
  case "./data/level1/small/m-165.png": return level1__small__m_165_png
  case "./data/level1/small/m-180.png": return level1__small__m_180_png
  case "./data/level1/small/m-30.png": return level1__small__m_30_png
  case "./data/level1/small/m-45.png": return level1__small__m_45_png
  case "./data/level1/small/m-60.png": return level1__small__m_60_png
  case "./data/level1/small/m-75.png": return level1__small__m_75_png
  case "./data/level1/small/m-90.png": return level1__small__m_90_png
  case "./data/level1/small/p-0.png": return level1__small__p_0_png
  case "./data/level1/small/p-105.png": return level1__small__p_105_png
  case "./data/level1/small/p-120.png": return level1__small__p_120_png
  case "./data/level1/small/p-135.png": return level1__small__p_135_png
  case "./data/level1/small/p-15.png": return level1__small__p_15_png
  case "./data/level1/small/p-150.png": return level1__small__p_150_png
  case "./data/level1/small/p-165.png": return level1__small__p_165_png
  case "./data/level1/small/p-30.png": return level1__small__p_30_png
  case "./data/level1/small/p-45.png": return level1__small__p_45_png
  case "./data/level1/small/p-60.png": return level1__small__p_60_png
  case "./data/level1/small/p-75.png": return level1__small__p_75_png
  case "./data/level1/small/p-90.png": return level1__small__p_90_png
  case "./data/level2/0-collision.png": return level2__0_collision_png
  case "./data/level2/0.png": return level2__0_png
  case "./data/level2/1-collision.png": return level2__1_collision_png
  case "./data/level2/1.png": return level2__1_png
  case "./data/level2/2-collision.png": return level2__2_collision_png
  case "./data/level2/2.png": return level2__2_png
  case "./data/level2/3-collision.png": return level2__3_collision_png
  case "./data/level2/3.png": return level2__3_png
  case "./data/level2/small/m-105.png": return level2__small__m_105_png
  case "./data/level2/small/m-120.png": return level2__small__m_120_png
  case "./data/level2/small/m-135.png": return level2__small__m_135_png
  case "./data/level2/small/m-15.png": return level2__small__m_15_png
  case "./data/level2/small/m-150.png": return level2__small__m_150_png
  case "./data/level2/small/m-165.png": return level2__small__m_165_png
  case "./data/level2/small/m-180.png": return level2__small__m_180_png
  case "./data/level2/small/m-30.png": return level2__small__m_30_png
  case "./data/level2/small/m-45.png": return level2__small__m_45_png
  case "./data/level2/small/m-60.png": return level2__small__m_60_png
  case "./data/level2/small/m-75.png": return level2__small__m_75_png
  case "./data/level2/small/m-90.png": return level2__small__m_90_png
  case "./data/level2/small/p-0.png": return level2__small__p_0_png
  case "./data/level2/small/p-105.png": return level2__small__p_105_png
  case "./data/level2/small/p-120.png": return level2__small__p_120_png
  case "./data/level2/small/p-135.png": return level2__small__p_135_png
  case "./data/level2/small/p-15.png": return level2__small__p_15_png
  case "./data/level2/small/p-150.png": return level2__small__p_150_png
  case "./data/level2/small/p-165.png": return level2__small__p_165_png
  case "./data/level2/small/p-30.png": return level2__small__p_30_png
  case "./data/level2/small/p-45.png": return level2__small__p_45_png
  case "./data/level2/small/p-60.png": return level2__small__p_60_png
  case "./data/level2/small/p-75.png": return level2__small__p_75_png
  case "./data/level2/small/p-90.png": return level2__small__p_90_png
  case "./data/level3/0-collision.png": return level3__0_collision_png
  case "./data/level3/0.png": return level3__0_png
  case "./data/misc/introduction.png": return misc__introduction_png
  case "./data/misc/loading.png": return misc__loading_png
  case "./data/misc/missing_clues.png": return misc__missing_clues_png
  case "./data/misc/parchment0.png": return misc__parchment0_png
  case "./data/misc/parchment1.png": return misc__parchment1_png
  case "./data/misc/parchment2.png": return misc__parchment2_png
  case "./data/misc/player_leftbottom.png": return misc__player_leftbottom_png
  case "./data/misc/player_lefttop.png": return misc__player_lefttop_png
  case "./data/misc/player_rightbottom.png": return misc__player_rightbottom_png
  case "./data/misc/player_righttop.png": return misc__player_righttop_png
  case "./data/misc/win_the_game.png": return misc__win_the_game_png
  default: panic("Could not find: "+filename)
  }
  return []byte{}
}
